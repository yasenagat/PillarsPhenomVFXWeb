package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddNote(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddNote: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var note u.ShotNote
	err = json.Unmarshal(data, &note)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddNote: json.Unmarshal(data, &note) failed!")
		return
	}
	// TODO 检查传入字段的有效性
	note.NoteCode = *u.GenerateCode(&userCode)
	note.UserCode = userCode

	err = postStorage.AddNote(&note)
	if err != nil {
		u.OutputJsonLog(w, 13, err.Error(), nil, "postAction.AddNote: postStorage.AddNote(&note) failed!")
		return
	}

	u.OutputJson(w, 0, "Add Note success.", note)
}

func QueryNotes(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryNotes: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	err = json.Unmarshal(data, &code)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryNotes: json.Unmarshal(data, &note) failed!")
		return
	}
	// TODO 检查传入字段的有效性

	result, err := postStorage.QueryNotes(&code)
	if result == nil || err != nil {
		u.OutputJsonLog(w, 13, "Query Notes failed!", nil, "postAction.QueryNotes: postStorage.QueryNotes(&note) failed!")
		return
	}

	u.OutputJson(w, 0, "Query success.", result)
}

// ----------------------------------------------------------------

// RECEVE: noteCode   RETURN: nil
//func DeleteNote(w http.ResponseWriter, r *http.Request) {
//	data, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		u.OutputJson(w, 1, "Read body failed!", nil)
//		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
//		return
//	}
//	var code string
//	json.Unmarshal(data, &code)
//	err = postStorage.DeleteNote(&code)
//	if err != nil {
//		u.OutputJson(w, 2, "Read body failed!", nil)
//		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
//		return
//	}
//	u.OutputJson(w, 0, "Delete success.", nil)
//}
