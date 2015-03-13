package postAction

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"os"

	"io"
	"io/ioutil"
	"net/http"
)

func LoadEdlFile(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		u.OutputJson(w, 1, "parse upload error!", nil)
		return
	}
	formData := r.MultipartForm
	files := formData.File["files"]
	if len(files) > 0 {
		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			u.OutputJson(w, 12, "open edl file error!", nil)
			return
		}
		out, err := os.Create("./upload/" + files[0].Filename)
		defer out.Close()
		if err != nil {
			u.OutputJson(w, 13, "create edl file failed!", nil)
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println("io failed")
			u.OutputJson(w, 14, "io copy edl file failed!", nil)
		}
		// 解析edl文件得到镜头的信息
		edlShots, err := u.ReadEdl(out.Name())
		if err != nil && err.Error() != "EOF" {
			msg := "解析文件错误:" + err.Error()
			u.OutputJson(w, 15, msg, nil)
			return
		}
		if len(edlShots) == 0 {
			u.OutputJson(w, 16, "edl not find short!", nil)
			return
		}
		projectCode := formData.Value["ProjectCode"][0]
		// 查询镜头关联素材的详细信息
		shots, err := postStorage.EdlShotsToShots(files[0].Filename, projectCode, edlShots)
		if shots == nil || err != nil {
			u.OutputJson(w, 17, "edl not find material!", nil)
			return
		}
		// 保存镜头信息
		err = postStorage.InsertMultipleShot(userCode, projectCode, shots)
		if err != nil {
			u.OutputJson(w, 18, err.Error(), nil)
			return
		}
		u.OutputJson(w, 0, "upload success!", shots)
		return
	}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "not find upload file!", nil)
}

// RECEVE: noteCode  RETURN: notes struct  (an Array)
func QueryShotByShotCode(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryShotByShotCode(&code)
	if err != nil {
		u.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	u.OutputJson(w, 0, "Deleteshot success.", result)
}

//shot struct
func AddShot(w http.ResponseWriter, r *http.Request) {
	//1:jie shou json zifuchuan
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	//2:  json  unmarshal()  get the shot object
	var shot u.Shot
	err = json.Unmarshal(data, &shot)
	if err != nil {
		u.OutputJson(w, 1, err.Error(), nil)
		return
	}
	name := "nShot"
	shot.ShotCode = *u.GenerateCode(&name)
	//3: insert into mysql
	err = postStorage.AddSingleShot(&shot)
	if err != nil {
		u.OutputJson(w, 1, err.Error(), nil)
		return
	}
	result, err := postStorage.QueryShotByShotCode(&shot.ShotCode)
	if err != nil {
		u.OutputJson(w, 1, err.Error(), nil)
		return
	}
	//4:return shot to client
	//content,_:= json.Marshal(shot)
	//这里缺少前后台通信规则！！按照自定义协议，打包协议结构，然后json编码为字符串，发往前端
	u.OutputJson(w, 0, "addshot success.", result)
}

////shotCode
func DeleteShot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteSingleShot(&code)
	if err != nil {
		u.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	u.OutputJson(w, 0, "Deleteshot success.", nil)
}

type ss struct {
	ShotCode string
	ShotName string
}

//shotCode ,shotName
func ModifyShotName(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var result ss
	json.Unmarshal(data, &result)
	err = postStorage.ModifyShotName(&result.ShotCode, &result.ShotName)
	if err != nil {
		u.OutputJson(w, 1, "ModifyShotName error", err)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	u.OutputJson(w, 0, "ModifyShotName success.", nil)
}

// RECEVE: sourceFile name   RETURN: notes struct  (an Array)
func QueryShotByProjectCode(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var name string
	json.Unmarshal(data, &name)
	result, err := postStorage.QueryShotByProjectCode(&name)
	if err != nil {
		u.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	u.OutputJson(w, 0, "Deleteshot success.", result)
}

/*      限定上传文件大小
type Sizer interface {
	Size() int64
}

//上传处理函数
func GetFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		msg := "获取上传文件错误:" + err.Error()
		fmt.Fprintf(w, msg)
		return
	}
	defer file.Close()
	if fileSizer, ok := file.(Sizer); ok {
		fileSize := fileSizer.Size() / (1024 * 1024)
		fmt.Println("上传文件的大小为: %d", fileSize, handler.Filename)
		if fileSize > 100 {
			msg := "获取上传文件错误:文件大小超出100M"
			fmt.Fprintf(w, msg)
			return
		}
	} else {
		msg := "获取上传文件错误:无法读取文件大小"
		fmt.Fprintf(w, msg)
		return
	}
	fmt.Fprintf(w, "ok")
}
*/
