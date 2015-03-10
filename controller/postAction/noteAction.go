package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//RECEVE: note struc  RETURN: note struct
func AddNewNote(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		return
	}
	var note utility.Note
	json.Unmarshal(data, &note)
	name := "nNote"
	note.NoteCode = *utility.GenerateCode(&name)
	err = postStorage.AddNote(&note)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	result, err := postStorage.QuerySingleNote(&note.NoteCode)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	utility.OutputJson(w, 0, "addAction success.", result)
}

// RECEVE: noteCode   RETURN: nil
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteNote(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		////pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Delete success.", nil)
}

// RECEVE: noteCode  RETURN: notes struct  (an Array)
func QueryNoteByShot(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryNoteByShot(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
