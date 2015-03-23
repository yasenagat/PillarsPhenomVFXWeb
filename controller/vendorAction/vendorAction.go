package vendorAction

import (
	s "PillarsPhenomVFXWeb/session"
	ps "PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
