package postAction

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"os"

	"io"
	"io/ioutil"
	"net/http"
)

func LoadEdlFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		msg := "获取上传文件错误:" + err.Error()
		utility.OutputJson(w, 1, msg, nil)
		return
	}
	defer file.Close()
	fmt.Println("---->3")
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 066)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Println("----->method:", r.Method)
	///////result 为镜头的信息
	//result, err := utility.AnalyseEdl(file)
	//if err != nil && err.Error() != "EOF" {
	//	msg := "解析文件错误:" + err.Error()
	//	utility.OutputJson(w, 2, msg, nil)
	//	return
	//}
	//if len(result) == 0 {
	//	utility.OutputJson(w, 3, "文件格式是否正确", nil)
	//	return
	//}
	//code := "ssssssssssssssssssss" //project code
	/**************数据库插入 about result********************/
	//err = postStorage.InsertMultipleShot(code, result)
	//if err != nil {
	//	utility.OutputJson(w, 4, err.Error(), nil)
	//	return
	//}
	////////////////////插入成功后该做什么处理？？？!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//utility.OutputJson(w, 0, "EDL FIle 插入成功.", nil)

}

////shot struct
func AddShot(w http.ResponseWriter, r *http.Request) {
	//1:jie shou json zifuchuan
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	//2:  json  unmarshal()  get the shot object
	var shot utility.Shot
	err = json.Unmarshal(data, &shot)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	name := "nShot"
	shot.ShotCode = *utility.GenerateCode(&name)
	//3: insert into mysql
	err = postStorage.AddSingleShot(&shot)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	result, err := postStorage.QueryShotByShotCode(&shot.ShotCode)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	//4:return shot to client
	//content,_:= json.Marshal(shot)
	//这里缺少前后台通信规则！！按照自定义协议，打包协议结构，然后json编码为字符串，发往前端
	utility.OutputJson(w, 0, "addshot success.", result)
}

////shotCode
func DeleteShot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteSingleShot(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Deleteshot success.", nil)
}

type ss struct {
	ShotCode string
	ShotName string
}

//shotCode ,shotName
func ModifyShotName(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var result ss
	json.Unmarshal(data, &result)
	err = postStorage.ModifyShotName(&result.ShotCode, &result.ShotName)
	if err != nil {
		utility.OutputJson(w, 1, "ModifyShotName error", err)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "ModifyShotName success.", nil)
}

// RECEVE: noteCode  RETURN: notes struct  (an Array)
func QueryShotByShotCode(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryShotByShotCode(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Deleteshot success.", result)
}

// RECEVE: sourceFile name   RETURN: notes struct  (an Array)
func QueryShotByProjectCode(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var name string
	json.Unmarshal(data, &name)
	result, err := postStorage.QueryShotByProjectCode(&name)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Deleteshot success.", result)
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
