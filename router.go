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
	http.HandleFunc("/post_shot_list", postAction.QueryShots)
	http.HandleFunc("/post_shot_que", postAction.QueryShotByShotCode)
	http.HandleFunc("/post_shot_upd", postAction.UpdateShot)
	http.HandleFunc("/post_shot_add", postAction.AddShot)
	http.HandleFunc("/post_shot_updshotname", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_demand_add", postAction.AddDemand)
	http.HandleFunc("/post_shot_demand_del", postAction.DeleteDemand)
	http.HandleFunc("/post_shot_demand_upd", postAction.UpdateDemand)
	http.HandleFunc("/post_shot_demand_que", postAction.QueryDemands)
	http.HandleFunc("/post_shot_note_add", postAction.AddNote)
	http.HandleFunc("/post_shot_note_que", postAction.QueryNotes)

	http.HandleFunc("/post_shot_folder", postAction.GetShotList)
	http.HandleFunc("/post_shot_folder_shots", postAction.QueryFolderShots)
	http.HandleFunc("/post_shot_folder_add", postAction.AddFolder)
	http.HandleFunc("/post_shot_folder_del", postAction.DeleteFolder)
	http.HandleFunc("/post_shot_folder_que", postAction.QueryFolder)
	http.HandleFunc("/post_shot_folder_upd", postAction.UpdateFolder)
	http.HandleFunc("/post_shot_folder_addfiles", postAction.AddFolderFiles)
	http.HandleFunc("/post_shot_folder_delfiles", postAction.DeleteFolderFiles)
	http.HandleFunc("/post_shot_folder_countfiles", postAction.CountFolderFiles)

	// ---------------------------- 待实现 ------------------------------
	http.HandleFunc("/post_shot_del", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_material_que", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_material_add", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_material_upd", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_material_del", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_version_que", postAction.ModifyShotName)

	//
	http.HandleFunc("/editoral_download_file_check", downloadAction.DownloadFileCheck)
	//http.HandleFunc("/.*", NotFound) // TODO 想实现未知路由地址访问的404页面跳转
}
