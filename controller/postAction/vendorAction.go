package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddVendor(w http.ResponseWriter, r *http.Request) {
	//data, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	utility.OutputJson(w, 1, "Read body failed!", nil)
	//	return
	//}
	//var vender utility.Vendor
	//json.Unmarshal(data, &vender)
	//name := "nVendor"
	//vender.VendorCode = *utility.GenerateCode(&name)
	//err = postStorage.AddVendor(&vender)
	//if err != nil {
	//	utility.OutputJson(w, 1, err.Error(), nil)
	//	return
	//}
	//result, err := postStorage.QuerySingleNote(&vender.VendorCode)
	//if err != nil {
	//	utility.OutputJson(w, 1, err.Error(), nil)
	//	return
	//}
	//utility.OutputJson(w, 0, "addAction success.", result)
}

/////////////////还要删除关联的镜头！！！！！！！！！！！
func DeleteVendor(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteVendor(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Delete success.", nil)
}

type Rauth struct {
	Code string
	Auth utility.VendorAuth
}

//////
func ModifyVendorAuth(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var rauth Rauth
	json.Unmarshal(data, &rauth)
	err = postStorage.ModifyVendorAuth(&rauth.Code, &rauth.Auth)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Modify success.", nil)
}

/* 查询外包商列表*/
func QueryVendorByPro(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryVendorByPro(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
