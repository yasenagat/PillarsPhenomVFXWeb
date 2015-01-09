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
	http.HandleFunc("/update", action.UpdateUserAction)
	http.HandleFunc("/query", action.QueryUserAction)
	http.HandleFunc("/userlist", action.UserListAction)

	http.HandleFunc("/addproject", action.AddProjectAction)
	http.HandleFunc("/delproject", action.DeleteProjectAction)
	http.HandleFunc("/updateproject", action.UpdateProjectAction)
	http.HandleFunc("/queryproject", action.QueryProjectAction)
	http.HandleFunc("/projectlist", action.ProjectListAction)
}
