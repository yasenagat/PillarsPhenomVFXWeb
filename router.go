package main

import (
	"PillarsPhenomVFXWeb/controller/downloadAction"
	"PillarsPhenomVFXWeb/controller/editoralAction"
	"PillarsPhenomVFXWeb/controller/loginAction"
	"PillarsPhenomVFXWeb/controller/postAction"
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
	http.HandleFunc("/project_find", projectAction.FindProjects)

	http.HandleFunc("/editoral_library", editoralAction.GetLibrarys)
	http.HandleFunc("/editoral_library_add", editoralAction.AddLibrary)
	http.HandleFunc("/editoral_library_materials", editoralAction.GetLibraryFileList)
	http.HandleFunc("/editoral_find_materials", editoralAction.FindMaterials)
	http.HandleFunc("/editoral_filetype", editoralAction.GetFiletypes)
	http.HandleFunc("/editoral_material", editoralAction.GetMaterialInfo)
	http.HandleFunc("/editoral_folder", editoralAction.GetFolders)
	http.HandleFunc("/editoral_folder_materials", editoralAction.QueryFolderMaterials)
	http.HandleFunc("/editoral_folder_add", editoralAction.AddFolder)
	http.HandleFunc("/editoral_folder_del", editoralAction.DeleteFolder)
	http.HandleFunc("/editoral_folder_que", editoralAction.QueryFolder)
	http.HandleFunc("/editoral_folder_upd", editoralAction.UpdateFolder)
	http.HandleFunc("/editoral_folder_addfiles", editoralAction.AddFolderFiles)
	http.HandleFunc("/editoral_folder_delfiles", editoralAction.DeleteFolderFiles)
	http.HandleFunc("/editoral_folder_countfiles", editoralAction.CountFolderFiles)
	http.HandleFunc("/editoral_download_file", downloadAction.DownloadFile)

	http.HandleFunc("/post_upload_edl", postAction.LoadEdlFile)
	// ---------------------------- 尚未测试 ------------------------------
	http.HandleFunc("/post_shot_que", postAction.QueryShotByShotCode)

	http.HandleFunc("/editoral_download_file_check", downloadAction.DownloadFileCheck)
	//http.HandleFunc("/.*", NotFound) // TODO 想实现未知路由地址访问的404页面跳转
}
