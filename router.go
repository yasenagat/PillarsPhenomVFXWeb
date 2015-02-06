package main

import (
	"PillarsPhenomVFXWeb/controller/loginAction"
	"PillarsPhenomVFXWeb/controller/projectAction"
	"PillarsPhenomVFXWeb/controller/userAction"
	"net/http"
)

func RouterBinding() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./pages"))))

	http.HandleFunc("/login_action", loginAction.Login)

	http.HandleFunc("/user_add", userAction.AddUser)
	http.HandleFunc("/user_del", userAction.DeleteUser)
	http.HandleFunc("/user_upd", userAction.UpdateUser)
	http.HandleFunc("/user_sel", userAction.QueryUser)
	http.HandleFunc("/user_list", userAction.UserList)

	http.HandleFunc("/project_add", projectAction.AddProject)
	http.HandleFunc("/project_del", projectAction.DeleteProject)
	http.HandleFunc("/project_upd", projectAction.UpdateProject)
	http.HandleFunc("/project_sel", projectAction.QueryProject)
	http.HandleFunc("/project_list", projectAction.ProjectList)
	http.HandleFunc("/project_load", projectAction.LoadProject)
	//http.HandleFunc("/.*", NotFound) // TODO 想实现未知路由地址访问的404页面跳转
}
