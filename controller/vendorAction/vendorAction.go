package vendorAction

import (
	s "PillarsPhenomVFXWeb/session"
	ps "PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

//查询外包商列表
func GetVendorProjectList(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "分包商")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	result, err := ps.GetVendorProject(&userCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 1, "Query failed!", nil, "companyAction.GetProjectList: postStorage.GetVendorProject(&userCode) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}

//分包商项目的镜头列表及需求列表信息
func QueryVendorProjectShots(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "分包商")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotVendorShots: ioutil.ReadAll(r.Body) failed!")
		return
	}

	type interim struct {
		VendorCode string
	}

	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotVendorShots: json.Unmarshal(data, &shots) failed!")
		return
	}
	if len(i.VendorCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.QueryShotVendorShots: Parameters Checked failed!")
		return
	}
	//镜头List
	result, err := ps.QueryVendorProjectShots(i.VendorCode, userCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 14, "Query shots failed!", nil, "postAction.QueryVendorProjectShots: postStorage.QueryShotVendorShots(VendorCode) failed!")
		return
	}

	type vendorShots struct {
		Shot    u.Shot
		Demands *[]u.ShotDemand
	}
	var vss []vendorShots
	//镜头的需求
	for _, s := range *result {
		var vs vendorShots
		vs.Shot = s
		ds, err := ps.QueryVendorDemands(&s.ShotCode)
		if err != nil || ds == nil {
			u.OutputJsonLog(w, 15, "Query shotDemands failed!", nil, "postAction.QueryDemands: postStorage.QueryDemands(&ShotCode) failed!")
			return
		}
		vs.Demands = ds
		vss = append(vss, vs)
	}

	u.OutputJsonLog(w, 0, "Query success.", vss, "")
}

// 判断文件是否存在: 存在返回true,不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func UploadDemo(w http.ResponseWriter, r *http.Request) {
	//flag, userCode := s.GetAuthorityCode(w, r, "制片")
	//if !flag {
	//	u.OutputJson(w, 404, "session error!", nil)
	//	return
	//}

	//err := r.ParseMultipartForm(32 << 20)
	//if err != nil {
	//	u.OutputJsonLog(w, 12, "parse upload error!", nil, "postAction.AddShotMaterial: r.ParseMultipartForm(32 << 20) failed!")
	//	return
	//}
	//formData := r.MultipartForm
	//var sm u.ShotMaterial

	//sm.ShotCode = formData.Value["ShotCode"][0]
	//sm.ProjectCode = formData.Value["ProjectCode"][0]
	//sm.MaterialType = formData.Value["MaterialType"][0]
	//sm.MaterialDetail = formData.Value["MaterialDetail"][0]
	//if len(sm.ShotCode) == 0 || len(sm.ProjectCode) == 0 || len(sm.MaterialType) == 0 || len(sm.MaterialDetail) == 0 {
	//	u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.AddShotMaterial: Parameter Checked failed!")
	//	return
	//}
	//files := formData.File["files"]
	//if len(files) > 0 {
	//	sm.MaterialName = files[0].Filename
	//	file, err := files[0].Open()
	//	defer file.Close()
	//	if err != nil {
	//		u.OutputJsonLog(w, 14, "Open upload file failed!", nil, "postAction.AddShotMaterial: Open upload file failed!")
	//		return
	//	}
	//	var path = "/home/pillars/Upload/material/" + sm.ProjectCode
	//	err = os.MkdirAll(path, 0777)
	//	if err != nil {
	//		u.OutputJsonLog(w, 15, "Create file path failed!", nil, "postAction.AddShotMaterial: Create file path failed!")
	//		return
	//	}
	//	createFile := path + "/" + sm.MaterialName
	//	if checkFileIsExist(createFile) { //如果文件存在
	//		u.OutputJsonLog(w, 202, "File Exist!", nil, "postAction.AddShotMaterial: File Exist!")
	//		return
	//	}
	//	out, err := os.OpenFile(createFile, os.O_CREATE|os.O_RDWR, 0777)
	//	if err != nil {
	//		u.OutputJsonLog(w, 16, "Create file failed!", nil, "postAction.AddShotMaterial: Create file failed!")
	//		return
	//	}
	//	defer out.Close()
	//	_, err = io.Copy(out, file)
	//	if err != nil {
	//		u.OutputJsonLog(w, 17, "Copy file failed!", nil, "postAction.AddShotMaterial: Copy file failed!")
	//		return
	//	}

	//	// TODO 文件上传,保存成功,是否需要调用C++对素材抓图及其他信息
	//	sm.MaterialCode = *u.GenerateCode(&userCode)
	//	sm.MaterialPath = out.Name()
	//	sm.UserCode = userCode
	//	err = postStorage.AddShotMaterial(&sm)
	//	if err != nil {
	//		u.OutputJsonLog(w, 18, err.Error(), nil, "postAction.AddShotMaterial: postStorage.AddShotMaterial(&ShotMaterial) failed!")
	//		return
	//	}

	//	u.OutputJsonLog(w, 0, "Upload success.", nil, "")
	//	return
	//}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "not find upload file!", nil)
}
