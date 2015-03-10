package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//RECEVE: reference struc  RETURN: note struct
func AddReference(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		return
	}
	var ref utility.Reference
	json.Unmarshal(data, &ref)
	name := "nReference"
	ref.ReferenceCode = *utility.GenerateCode(&name)
	err = postStorage.AddReference(&ref)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	result, err := postStorage.QuerySingleReference(&ref.ReferenceCode)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	utility.OutputJson(w, 0, "addAction success.", result)
}

// RECEVE: refCode   RETURN: nil
func DeleteReference(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteReference(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Delete success.", nil)
}

// RECEVE: refCode  RETURN: refs struct  (an Array)
func QueryReferenceByShot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryReferenceByShot(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
func QueryRefDetail(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryRefDetail(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
