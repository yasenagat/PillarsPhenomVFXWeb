package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func AddShotMaterial(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 1, "session error!", nil)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		u.OutputJsonLog(w, 12, "parse upload error!", nil, "postAction.AddShotMaterial: r.ParseMultipartForm(32 << 20) failed!")
		return
	}
	formData := r.MultipartForm
	projectCode := formData.Value["ProjectCode"][0]
	if len(projectCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.AddShotMaterial: Parameter Checked failed!")
		return
	}
	files := formData.File["files"]
	if len(files) > 0 {
		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			u.OutputJsonLog(w, 14, "Open upload file failed!", nil, "postAction.AddShotMaterial: Open upload file failed!")
			return
		}
		var path = "./upload/material/" + projectCode + "/" + files[0].Filename
		out, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			u.OutputJsonLog(w, 15, "Create file failed!", nil, "postAction.AddShotMaterial: Create file failed!")
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			u.OutputJsonLog(w, 16, "Copy file failed!", nil, "postAction.AddShotMaterial: Copy file failed!")
			return
		}
		fmt.Println("--------->")
		return
	}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "not find upload file!", nil)
}

func DeleteShotMaterial(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteDemand: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteDemand: json.Unmarshal(data, &demand) failed!")
		return
	}
	if len(demand.DemandCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.DeleteDemand: Parameters Checked failed!")
		return
	}
	demand.UserCode = userCode

	err = postStorage.DeleteDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.DeleteDemand: postStorage.DeleteDemand(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "Delete success.", nil)
}

func UpdateShotMaterial(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.UpdateDemand: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.UpdateDemand: json.Unmarshal(data, &demand) failed!")
		return
	}
	if len(demand.DemandCode) == 0 || (len(demand.DemandDetail) == 0 && len(demand.Picture) == 0) {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.UpdateDemand: Parameter Checked failed!")
		return
	}
	demand.UserCode = userCode

	err = postStorage.UpdateDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.UpdateDemand: postStorage.UpdateDemand(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "Update success.", nil)
}

func QueryShotMaterials(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryDemands: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryDemands: json.Unmarshal(data, &ShotCode) failed!")
		return
	}
	if len(i.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameter ShotCode failed!", nil, "postAction.QueryDemands: Parameter ShotCode failed!")
		return
	}

	result, err := postStorage.QueryDemands(&i.ShotCode)
	if result == nil || err != nil {
		u.OutputJsonLog(w, 14, "Query ShotDemands failed!", nil, "postAction.QueryDemands: postStorage.QueryDemands(&ShotCode) failed!")
		return
	}

	u.OutputJson(w, 0, "Query success.", result)
}
