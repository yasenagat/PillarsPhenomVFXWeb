package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertLibrary(l *utility.Library) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO library (library_code, library_name, library_path, encoded_path, dpx_path, mov_path, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(l.LibraryCode, l.LibraryName, l.LibraryPath, l.EncodedPath, l.DpxPath, l.MovPath, l.UserCode, l.ProjectCode, l.Status)
	if err != nil {
		return false, err
	}

	return true, err
}
