package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

// Note struct
func AddNote(note *utility.Note) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO `note`(note_code,shot_code,content,thumbnail,status) VALUE(?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(note.NoteCode, note.ShotCode, note.Content, note.Image, note.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteNote(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `note` SET status=1 WHERE note_code=? ")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(code)
	if err != nil {
		return err
	}
	return nil
}
func QuerySingleNote(code *string) (*utility.Note, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT note_code,shot_code,content,thumbnail,status FROM note WHERE note_code=?  AND status=0")
	defer stmt.Close()
	var note utility.Note
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&note.NoteCode, &note.ShotCode, &note.Content, &note.Image, &note.Status)
	if err != nil {
		//fmt.Println("the note result scan appere one error")
		return nil, err
	}
	return &note, nil
}
func QueryNoteByShot(code *string) ([]utility.Note, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT note_code,shot_code,content,thumbnail,status FROM note WHERE shot_code=?  AND status=0")
	defer stmt.Close()
	var notes []utility.Note
	if err != nil {
		return notes, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return notes, err
	}
	if result.Next() {
		var note utility.Note
		err = result.Scan(&note.NoteCode, &note.ShotCode, &note.Content, &note.Image, &note.Status)
		if err != nil {
			return notes, err
		}
		notes = append(notes, note)

	}
	return notes, nil
}
