package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddDemand(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddDemand: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddDemand: json.Unmarshal(data, &demand) failed!")
		return
	}
	// TODO 检查传入字段的有效性
	demand.DemandCode = *u.GenerateCode(&userCode)
	demand.UserCode = userCode

	err = postStorage.AddDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 13, err.Error(), nil, "postAction.AddDemand: postStorage.AddDemand(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "addAction success.", demand)
}

func DeleteDemand(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteRequment: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteRequment: json.Unmarshal(data, &demand) failed!")
		return
	}
	// TODO 检查传入字段的有效性
	demand.UserCode = userCode

	err = postStorage.DeleteDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 13, err.Error(), nil, "postAction.DeleteRequment: postStorage.DeleteRequment(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "Delete success.", nil)
}

func UpdateDemand(w http.ResponseWriter, r *http.Request) {
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
	// TODO 检查传入字段的有效性
	demand.UserCode = userCode

	err = postStorage.UpdateDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 13, err.Error(), nil, "postAction.UpdateDemand: postStorage.UpdateDemand(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "Delete success.", nil)
}

func QueryDemands(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteRequment: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	err = json.Unmarshal(data, &code)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteRequment: json.Unmarshal(data, &demand) failed!")
		return
	}
	// TODO 检查传入字段的有效性

	result, err := postStorage.QueryDemands(&code)
	if result == nil || err != nil {
		u.OutputJsonLog(w, 13, "Query ShotDemands failed!", nil, "postAction.QueryDemands: postStorage.QueryReqByShot(&code) failed!")
		return
	}

	u.OutputJson(w, 0, "Query success.", result)
}