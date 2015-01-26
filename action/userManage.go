package action

import (
	"PillarsPhenomVFXWeb/mysqlStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func AddUserAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olen := len(r.Form["Email"]) + len(r.Form["UserName"]) + len(r.Form["Phone"]) + len(r.Form["UserAuthority"]) + len(r.Form["FilePath"])
	if olen != 5 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	if len(r.Form["UserName"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter UserName", nil)
		return
	}

	if len(r.Form["Phone"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter Phone", nil)
		return
	}

	if len(r.Form["UserAuthority"][0]) == 0 {
		u.OutputJson(w, 15, "Error parameter UserAuthority", nil)
		return
	}

	if len(r.Form["FilePath"][0]) == 0 {
		u.OutputJson(w, 16, "Error parameter FilePath", nil)
		return
	}

	temp := "insert"
	userCode := u.GenerateCode(&temp)
	picture := u.GenerateCode(&temp)
	user := u.User{
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
		u.OutputJson(w, 17, "Inert into user failed!", nil)
		return
	}

	UserListAction(w, r)
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
	list.Execute(w, userList)
}

func DeleteUserAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olen := len(r.Form["Email"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	email := r.Form["Email"][0]
	result, _ := mysqlStorage.DeleteUserByEmail(&email)
	if result == false {
		u.OutputJson(w, 13, "Delete user failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Delete user succeed!", nil)
}

func QueryUserAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olen := len(r.Form["Email"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	email := r.Form["Email"][0]
	user, err := mysqlStorage.QueryUserByEmail(&email)
	if err != nil {
		u.OutputJson(w, 13, "Query user failed!", nil)
		return
	}

	result, _ := json.Marshal(user)
	fmt.Fprintf(w, string(result))
}

func UpdateUserAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olen := len(r.Form["Email"]) + len(r.Form["UserName"]) + len(r.Form["Phone"]) + len(r.Form["UserAuthority"]) + len(r.Form["FilePath"])
	if olen != 5 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	if len(r.Form["UserName"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter UserName", nil)
		return
	}

	if len(r.Form["Phone"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter Phone", nil)
		return
	}

	if len(r.Form["UserAuthority"][0]) == 0 {
		u.OutputJson(w, 15, "Error parameter UserAuthority", nil)
		return
	}

	if len(r.Form["FilePath"][0]) == 0 {
		u.OutputJson(w, 16, "Error parameter FilePath", nil)
		return
	}

	user := u.User{
		DisplayName:   r.Form["UserName"][0],
		Email:         r.Form["Email"][0],
		Phone:         r.Form["Phone"][0],
		UserAuthority: r.Form["UserAuthority"][0],
		FilePath:      r.Form["FilePath"][0],
	}
	result, _ := mysqlStorage.UpdateUserByEmail(&user)
	if result == false {
		u.OutputJson(w, 17, "Update user failed!", nil)
		return
	}

	UserListAction(w, r)
}
