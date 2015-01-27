package main

import (
	"PillarsPhenomVFXWeb/action"
	"net/http"
)

func RouterBinding() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./pages"))))

	http.HandleFunc("/login_action", action.LoginAction)

	http.HandleFunc("/user_add", action.AddUserAction)
	http.HandleFunc("/user_del", action.DeleteUserAction)
	http.HandleFunc("/user_upd", action.UpdateUserAction)
	http.HandleFunc("/user_sel", action.QueryUserAction)
	http.HandleFunc("/user_list", action.UserListAction)

	http.HandleFunc("/project_add", action.AddProjectAction)
	http.HandleFunc("/project_del", action.DeleteProjectAction)
	http.HandleFunc("/project_upd", action.UpdateProjectAction)
	http.HandleFunc("/project_sel", action.QueryProjectAction)
	http.HandleFunc("/project_list", action.ProjectListAction)
}
