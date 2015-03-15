package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddNote(n *utility.ShotNote) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_note (note_code, shot_code, project_code, picture, note_detail, note_type, note_verson, user_code, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(n.NoteCode, n.ShotCode, n.ProjectCode, n.Picture, n.NoteDetail, n.NoteType, n.NoteVerson, n.UserCode, n.Status)
	if err != nil {
		return err
	}
	return nil
}

func QueryNotes(code *string) (*[]utility.ShotNote, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT note_code, picture, note_detail, note_type, note_verson FROM shot_note WHERE status = 0 AND shot_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(code)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var notes []utility.ShotNote
	for result.Next() {
		var note utility.ShotNote
		err = result.Scan(&note.NoteCode, &note.Picture, &note.NoteDetail, &note.NoteType, &note.NoteVerson)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return &notes, nil
}

// -------------------------------------------------------------------

//
func QuerySingleNote(code *string) (*utility.ShotNote, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT note_code,shot_code,note_detail,picture,note_verson FROM shot_note WHERE status = 0 AND note_code = ?")
	defer stmt.Close()
	var note utility.ShotNote
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&note.NoteCode, &note.ShotCode, &note.NoteDetail, &note.Picture, &note.NoteVerson)
	if err != nil {
		//fmt.Println("the note result scan appere one error")
		return nil, err
	}
	return &note, nil
}

//func DeleteNote(code *string) error {
//	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `note` SET status=1 WHERE note_code=? ")
//	defer stmt.Close()
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(code)
//	if err != nil {
//		return err
//	}
//	return nil
//}
