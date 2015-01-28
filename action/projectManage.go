package action

import (
	"PillarsPhenomVFXWeb/mysqlStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func AddProjectAction(w http.ResponseWriter, r *http.Request) {
	if !checkAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectName"]) + len(r.Form["ProjectDetail"]) + len(r.Form["ProjectLeader"])
	if olen != 3 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectName"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectName", nil)
		return
	}

	if len(r.Form["ProjectDetail"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter ProjectDetail", nil)
		return
	}

	if len(r.Form["ProjectLeader"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter ProjectLeader", nil)
		return
	}

	temp := "insert"
	projectCode := u.GenerateCode(&temp)
	picture := u.GenerateCode(&temp)
	project := u.Project{
		ProjectCode:   *projectCode,
		ProjectName:   r.Form["ProjectName"][0],
		Picture:       *picture,
		ProjectDetail: r.Form["ProjectDetail"][0],
		ProjectLeader: r.Form["ProjectLeader"][0],
		Status:        0,
	}
	result, _ := mysqlStorage.InsertIntoProject(&project)
	if result == false {
		u.OutputJson(w, 15, "Insert into project failed!", nil)
		return
	}

	ProjectListAction(w, r)
}

func DeleteProjectAction(w http.ResponseWriter, r *http.Request) {
	if !checkAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}

	code := r.Form["ProjectCode"][0]
	result, _ := mysqlStorage.DeleteProjectByProjectCode(&code)
	if result == false {
		u.OutputJson(w, 13, "Delete project failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Delete project succeed!", nil)
}

func UpdateProjectAction(w http.ResponseWriter, r *http.Request) {
	if !checkAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["ProjectName"]) + len(r.Form["ProjectDetail"]) + len(r.Form["ProjectLeader"])
	if olen != 4 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}

	if len(r.Form["ProjectName"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter ProjectName", nil)
		return
	}

	if len(r.Form["ProjectDetail"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter ProjectDetail", nil)
		return
	}

	if len(r.Form["ProjectLeader"][0]) == 0 {
		u.OutputJson(w, 15, "Error parameter ProjectLeader", nil)
		return
	}

	project := u.Project{
		ProjectCode:   r.Form["ProjectCode"][0],
		ProjectName:   r.Form["ProjectName"][0],
		ProjectDetail: r.Form["ProjectDetail"][0],
		ProjectLeader: r.Form["ProjectLeader"][0],
	}
	result, _ := mysqlStorage.UpdateProjectByProjectCode(&project)
	if result == false {
		u.OutputJson(w, 16, "Update project failed!", nil)
		return
	}

	ProjectListAction(w, r)
}

func QueryProjectAction(w http.ResponseWriter, r *http.Request) {
	if !checkAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}
	code := r.Form["ProjectCode"][0]
	project, err := mysqlStorage.QueryProjectByProjectCode(&code)
	if err != nil {
		u.OutputJson(w, 13, "Query project failed!", nil)
		return
	}

	result, _ := json.Marshal(project)
	fmt.Fprintf(w, string(result))
}

func ProjectListAction(w http.ResponseWriter, r *http.Request) {
	if !checkAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	list, err := template.ParseFiles("pages/projectlist.gtpl")
	if err != nil {
		panic(err.Error())
	}
	projectList, err := mysqlStorage.QueryProjectList()
	if err != nil {
		panic(err.Error())
	}
	list.Execute(w, projectList)
}
