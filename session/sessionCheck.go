package session

import (
	us "PillarsPhenomVFXWeb/storage/userStorage"
	"net/http"
)

/*
	从session中获取当前登陆用户的UserCode
	在session中取不到userCode返回false
*/
func GetSessionUserCode(w http.ResponseWriter, r *http.Request) (bool, string) {
	userSession := GlobalSessions.SessionStart(w, r)
	userCode := userSession.Get("userCode")
	if userCode == nil || userCode == "" {
		return false, ""
	}
	return true, string(userCode.(string))
}

/*
	session和权限的验证
*/
func CheckAuthority(w http.ResponseWriter, r *http.Request, authority string) bool {
	flag, s_code := GetSessionUserCode(w, r)
	if flag == false || s_code == "" {
		return false
	}

	rs, _ := us.GetUserAuthority(&s_code)
	if *rs != "admin" && *rs != authority {
		return false
	}
	return true
}
