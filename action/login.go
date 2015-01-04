package action

import (
	"PillarsPhenomVFXWeb/mysqlStorage"
	"PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LoginAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage // 返回json构建
	r.ParseForm()                           //form 数据解析

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["Password"]) +
		len(r.Form["Email"])

	if olen != 2 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["Password"][0]) == 0 {
		backMessage.FeedbackCode = 11
		backMessage.FeedbackText = "Error parameter Password"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}
	if len(r.Form["Email"][0]) == 0 {
		backMessage.FeedbackCode = 13
		backMessage.FeedbackText = "Error parameter Email"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}
	e := r.Form["Email"][0]
	p := r.Form["Password"][0]
	//fmt.Println("login user password:" + passwordMd5)
	userCode, err := mysqlStorage.CheckEmailAndPassword(&(e), &(p))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if *userCode == "" {
		backMessage.FeedbackCode = 1
		backMessage.FeedbackText = "Login failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))

	} else {
		fmt.Println(*userCode)
		userSession := session.GlobalSessions.SessionStart(w, r)
		userSession.Set("userCode", *userCode)
		userSession.Set("errorTimes", 0)
		userSession.Set("loginTime", time.Now().Unix())
		userSession.Set("lastAction", time.Now().Unix())
		backMessage.FeedbackCode = 0
		backMessage.FeedbackText = "Login succeed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	}
}
