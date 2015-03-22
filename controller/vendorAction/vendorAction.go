package vendorAction

import (
	s "PillarsPhenomVFXWeb/session"
	ps "PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
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
