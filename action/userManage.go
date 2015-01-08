package action

import (
	"PillarsPhenomVFXWeb/mysqlStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func AddUserAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage // 返回json构建
	r.ParseForm()                           //form 数据解析

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["Email"]) + len(r.Form["UserName"]) + len(r.Form["Phone"]) +
		len(r.Form["UserAuthority"]) + len(r.Form["FilePath"])

	if olen != 5 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
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

	if len(r.Form["UserName"][0]) == 0 {
		backMessage.FeedbackCode = 14
		backMessage.FeedbackText = "Error parameter UserName"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["Phone"][0]) == 0 {
		backMessage.FeedbackCode = 15
		backMessage.FeedbackText = "Error parameter Phone"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["UserAuthority"][0]) == 0 {
		backMessage.FeedbackCode = 16
		backMessage.FeedbackText = "Error parameter UserAuthority"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["FilePath"][0]) == 0 {
		backMessage.FeedbackCode = 17
		backMessage.FeedbackText = "Error parameter FilePath"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	temp := "insert"
	userCode := utility.GenerateCode(&temp)
	picture := utility.GenerateCode(&temp)
	user := utility.User{
		UserCode:      *userCode,
		Password:      "e10adc3949ba59abbe56e057f20f883e", // 默认为md5(123456, 32)
		DisplayName:   r.Form["UserName"][0],
		Picture:       *picture,
		Email:         r.Form["Email"][0],
		Phone:         r.Form["Phone"][0],
		UserAuthority: r.Form["UserAuthority"][0],
		FilePath:      r.Form["FilePath"][0],
		Status:        0,
	}
	result, _ := mysqlStorage.InsertIntoUser(&user)
	if result == false {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Inert into user failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		UserListAction(w, r)
		//backMessage.FeedbackCode = 0
		//backMessage.FeedbackText = "Save userinfo succeed!"
		//feedMessage, _ := json.Marshal(backMessage)
		//fmt.Fprintf(w, string(feedMessage))
	}

}

func UserListAction(w http.ResponseWriter, r *http.Request) {
	list, err := template.ParseFiles("pages/userlist.gtpl")
	if err != nil {
		panic(err.Error())
	}
	userList, err := mysqlStorage.QueryUserList()
	if err != nil {
		panic(err.Error())
	}
	//utility.Dump(userList)
	list.Execute(w, userList)
}

func DeleteUserAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage
	r.ParseForm()

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["Email"])

	if olen != 1 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
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

	email := r.Form["Email"][0]
	result, _ := mysqlStorage.DeleteUserByEmail(&email)
	if result == false {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Delete user failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		backMessage.FeedbackCode = 0
		backMessage.FeedbackText = "Delete user succeed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	}
}

func QueryUserAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage
	r.ParseForm()
	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["Email"])
	if olen != 1 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
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
	email := r.Form["Email"][0]
	user, err := mysqlStorage.QueryUserByEmail(&email)
	if err != nil {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Query user failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		feedMessage, _ := json.Marshal(user)
		fmt.Fprintf(w, string(feedMessage))
	}
}

func UpdateUserAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage
	r.ParseForm()

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["Email"]) + len(r.Form["UserName"]) + len(r.Form["Phone"]) +
		len(r.Form["UserAuthority"]) + len(r.Form["FilePath"])
	if olen != 5 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
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

	if len(r.Form["UserName"][0]) == 0 {
		backMessage.FeedbackCode = 14
		backMessage.FeedbackText = "Error parameter UserName"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["Phone"][0]) == 0 {
		backMessage.FeedbackCode = 15
		backMessage.FeedbackText = "Error parameter Phone"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["UserAuthority"][0]) == 0 {
		backMessage.FeedbackCode = 16
		backMessage.FeedbackText = "Error parameter UserAuthority"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["FilePath"][0]) == 0 {
		backMessage.FeedbackCode = 17
		backMessage.FeedbackText = "Error parameter FilePath"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	user := utility.User{
		DisplayName:   r.Form["UserName"][0],
		Email:         r.Form["Email"][0],
		Phone:         r.Form["Phone"][0],
		UserAuthority: r.Form["UserAuthority"][0],
		FilePath:      r.Form["FilePath"][0],
	}
	result, _ := mysqlStorage.UpdateUserByEmail(&user)
	if result == false {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Update user failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		UserListAction(w, r)
	}
}
