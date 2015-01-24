package action

import (
	"PillarsPhenomVFXWeb/mysqlStorage"
	"PillarsPhenomVFXWeb/session"
	u "PillarsPhenomVFXWeb/utility"
	"net/http"
	"time"
)

func LoginAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olen := len(r.Form["Password"]) + len(r.Form["Email"])
	if olen != 2 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Password"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Password", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter Email", nil)
		return
	}

	e := r.Form["Email"][0]
	p := r.Form["Password"][0]
	userCode, err := mysqlStorage.CheckEmailAndPassword(&(e), &(p))
	if err != nil {
		u.OutputJson(w, 14, "Email or Password wrong!", nil)
		return
	}

	if *userCode == "" {
		u.OutputJson(w, 15, "Login failed!", nil)
	}

	userSession := session.GlobalSessions.SessionStart(w, r)
	userSession.Set("userCode", *userCode)
	userSession.Set("errorTimes", 0)
	userSession.Set("loginTime", time.Now().Unix())
	userSession.Set("lastAction", time.Now().Unix())

	u.OutputJson(w, 0, "Login succeed!", nil)
}
