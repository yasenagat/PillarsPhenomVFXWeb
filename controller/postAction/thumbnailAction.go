package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddThumbnail(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		return
	}
	var image utility.Thumbnail
	json.Unmarshal(data, &image)
	name := "nThumbnail"
	image.ThumbnailCode = *utility.GenerateCode(&name)
	err = postStorage.AddThumbnail(&image)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	result, err := postStorage.QuerySingleImage(&image.ThumbnailCode)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	utility.OutputJson(w, 0, "addAction success.", result)
}

func DeleteThumbnail(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteThumbnail(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Delete success.", nil)
}
func QueryThumbnailByShotCode(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//	pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryThumbnailByShotCode(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
