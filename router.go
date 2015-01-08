package main

import (
	"PillarsPhenomVFXWeb/action"
	"net/http"
)

func RouterBinding() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./pages"))))

	http.HandleFunc("/login_action", action.LoginAction)
	http.HandleFunc("/add", action.AddUserAction)
	http.HandleFunc("/del", action.DeleteUserAction)
	http.HandleFunc("/userlist", action.UserListAction)
	http.HandleFunc("/query", action.QueryUserAction)
	http.HandleFunc("/update", action.UpdateUserAction)
}
