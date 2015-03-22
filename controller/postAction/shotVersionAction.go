package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddShotDemo(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddShotDemo: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var version u.ShotVersion
	json.Unmarshal(data, &version)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddShotDemo: json.Unmarshal(data, &version) failed!")
		return
	}
	if len(version.ProjectCode) == 0 || len(version.VendorUser) == 0 || len(version.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddShotDemo: Parameters Checked failed!")
		return
	}
	//获取版本号
	num, err := postStorage.GetShotVersionNum(&version)
	if num == nil || err != nil {
		u.OutputJsonLog(w, 14, "Get shot version number failed!", nil, "postAction.AddShotDemo: postStorage.GetShotVersionNum(&version) failed!")
		return
	}
	version.VersionNum = *num + 1
	version.VersionCode = *u.GenerateCode(&userCode)
	err = postStorage.AddShotDemo(&version)
	if err != nil {
		u.OutputJsonLog(w, 15, err.Error(), nil, "postAction.AddShotDemo: postStorage.AddShotDemo(&version) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Add success.", nil, "")
}

func QueryShotVersion(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotVersion: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var version u.ShotVersion
	json.Unmarshal(data, &version)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotVersion: json.Unmarshal(data, &version) failed!")
		return
	}
	if len(version.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.QueryShotVersion: Parameters Checked failed!")
		return
	}
	result, err := postStorage.QueryShotVersion(&version.ShotCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 15, "Query failed!", nil, "postAction.QueryShotVersion: postStorage.QueryShotVersion(&ShotCode) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", nil, "")
}

//func DeleteVersion(w http.ResponseWriter, r *http.Request) {
//	data, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		u.OutputJson(w, 1, "Read body failed!", nil)
//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
//		return
//	}
//	var code string
//	json.Unmarshal(data, &code)
//	err = postStorage.DeleteVersion(&code)
//	if err != nil {
//		u.OutputJson(w, 2, "Read body failed!", nil)
//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
//		return
//	}
//	u.OutputJson(w, 0, "Delete success.", nil)
//}
