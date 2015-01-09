package action

import (
	"PillarsPhenomVFXWeb/mysqlStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func AddProjectAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage // 返回json构建
	r.ParseForm()                           //form 数据解析

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["ProjectName"]) + len(r.Form["ProjectDescribe"]) +
		len(r.Form["ProjectLeader"])

	if olen != 3 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectName"][0]) == 0 {
		backMessage.FeedbackCode = 14
		backMessage.FeedbackText = "Error parameter ProjectName"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectDescribe"][0]) == 0 {
		backMessage.FeedbackCode = 15
		backMessage.FeedbackText = "Error parameter ProjectDescribe"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectLeader"][0]) == 0 {
		backMessage.FeedbackCode = 16
		backMessage.FeedbackText = "Error parameter ProjectLeader"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	temp := "insert"
	projectCode := utility.GenerateCode(&temp)
	picture := utility.GenerateCode(&temp)
	project := utility.Project{
		ProjectCode:     *projectCode,
		ProjectName:     r.Form["ProjectName"][0],
		Picture:         *picture,
		ProjectDescribe: r.Form["ProjectDescribe"][0],
		ProjectLeader:   r.Form["ProjectLeader"][0],
		Status:          0,
	}
	result, _ := mysqlStorage.InsertIntoProject(&project)
	if result == false {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Inert into project failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		ProjectListAction(w, r)
	}
}

func DeleteProjectAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage
	r.ParseForm()

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["ProjectCode"])

	if olen != 1 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		backMessage.FeedbackCode = 13
		backMessage.FeedbackText = "Error parameter ProjectCode"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	code := r.Form["ProjectCode"][0]
	result, _ := mysqlStorage.DeleteProjectByProjectCode(&code)
	if result == false {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Delete project failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		backMessage.FeedbackCode = 0
		backMessage.FeedbackText = "Delete project succeed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	}
}

func UpdateProjectAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage
	r.ParseForm()

	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["ProjectCode"]) + len(r.Form["ProjectName"]) +
		len(r.Form["ProjectDescribe"]) + len(r.Form["ProjectLeader"])
	if olen != 4 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		backMessage.FeedbackCode = 13
		backMessage.FeedbackText = "Error parameter ProjectCode"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectName"][0]) == 0 {
		backMessage.FeedbackCode = 14
		backMessage.FeedbackText = "Error parameter ProjectName"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectDescribe"][0]) == 0 {
		backMessage.FeedbackCode = 15
		backMessage.FeedbackText = "Error parameter ProjectDescribe"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	if len(r.Form["ProjectLeader"][0]) == 0 {
		backMessage.FeedbackCode = 16
		backMessage.FeedbackText = "Error parameter ProjectLeader"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}

	project := utility.Project{
		ProjectCode:     r.Form["ProjectCode"][0],
		ProjectName:     r.Form["ProjectName"][0],
		ProjectDescribe: r.Form["ProjectDescribe"][0],
		ProjectLeader:   r.Form["ProjectLeader"][0],
	}
	result, _ := mysqlStorage.UpdateProjectByProjectCode(&project)
	if result == false {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Update project failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		ProjectListAction(w, r)
	}
}

func QueryProjectAction(w http.ResponseWriter, r *http.Request) {
	var backMessage utility.FeedbackMessage
	r.ParseForm()
	//段落1： 判断参数个数是否正常。
	olen := len(r.Form["ProjectCode"])
	if olen != 1 {
		backMessage.FeedbackCode = 10
		backMessage.FeedbackText = "Error parameter format"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}
	if len(r.Form["ProjectCode"][0]) == 0 {
		backMessage.FeedbackCode = 13
		backMessage.FeedbackText = "Error parameter ProjectCode"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
		return
	}
	code := r.Form["ProjectCode"][0]
	project, err := mysqlStorage.QueryProjectByProjectCode(&code)
	if err != nil {
		backMessage.FeedbackCode = 2
		backMessage.FeedbackText = "Query project failed!"
		feedMessage, _ := json.Marshal(backMessage)
		fmt.Fprintf(w, string(feedMessage))
	} else {
		feedMessage, _ := json.Marshal(project)
		fmt.Fprintf(w, string(feedMessage))
	}
}

func ProjectListAction(w http.ResponseWriter, r *http.Request) {
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
