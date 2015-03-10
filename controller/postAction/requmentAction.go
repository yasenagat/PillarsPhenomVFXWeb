package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddRequment(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		return
	}
	var requment utility.Requment
	json.Unmarshal(data, &requment)
	name := "nRequment"
	requment.RequmentCode = *utility.GenerateCode(&name)
	err = postStorage.AddRequment(&requment)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	result, err := postStorage.QuerySingleRequment(&requment.RequmentCode)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	utility.OutputJson(w, 0, "addAction success.", result)
}

func DeleteRequment(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteRequment(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Delete success.", nil)
}
func QueryReqByShot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryReqByShot(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
