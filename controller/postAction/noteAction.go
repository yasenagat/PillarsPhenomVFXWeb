package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//RECEVE: note struc  RETURN: note struct
func AddNote(w http.ResponseWriter, r *http.Request) {
	//flag, userCode := s.GetAuthorityCode(w, r, "制片")
	//if !flag {
	//	http.Redirect(w, r, "/404.html", http.StatusFound)
	//	return
	//}

	//data, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddDemand: ioutil.ReadAll(r.Body) failed!")
	//	return
	//}
	//var note u.Note
	//err = json.Unmarshal(data, &note)
	//if err != nil {
	//	u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddDemand: json.Unmarshal(data, &demand) failed!")
	//	return
	//}
	//// TODO 检查传入字段的有效性
	//note.NoteCode = *u.GenerateCode(&userCode)
	//note.UserCode = userCode

	//err = postStorage.AddDemand(&demand)
	//if err != nil {
	//	u.OutputJsonLog(w, 13, err.Error(), nil, "postAction.AddDemand: postStorage.AddDemand(&demand) failed!")
	//	return
	//}

	//u.OutputJson(w, 0, "addAction success.", demand)

	//var note u.Note
	//json.Unmarshal(data, &note)
	//name := "nNote"
	//note.NoteCode = *u.GenerateCode(&name)
	//err = postStorage.AddNote(&note)
	//if err != nil {
	//	u.OutputJson(w, 1, err.Error(), nil)
	//	return
	//}
	//result, err := postStorage.QuerySingleNote(&note.NoteCode)
	//if err != nil {
	//	u.OutputJson(w, 1, err.Error(), nil)
	//	return
	//}
	//u.OutputJson(w, 0, "addAction success.", result)
}

// RECEVE: noteCode  RETURN: notes struct  (an Array)
func QueryNoteByShot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryNoteByShot(&code)
	if err != nil {
		u.OutputJson(w, 2, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	u.OutputJson(w, 0, "Query success.", result)
}

// ----------------------------------------------------------------

// RECEVE: noteCode   RETURN: nil
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteNote(&code)
	if err != nil {
		u.OutputJson(w, 2, "Read body failed!", nil)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	u.OutputJson(w, 0, "Delete success.", nil)
}
