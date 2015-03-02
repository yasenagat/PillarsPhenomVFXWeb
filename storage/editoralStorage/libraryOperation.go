package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertLibrary(l *utility.Library) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO library (library_code, library_name, library_path, dpx_path, jpg_path, mov_path, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(l.LibraryCode, l.LibraryName, l.LibraryPath, l.DpxPath, l.JpgPath, l.MovPath, l.UserCode, l.ProjectCode, l.Status)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}

	return true, err
}

func QueryLibraryByLibraryCode(code *string) (*utility.Library, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT library_code, library_name, library_path, dpx_path, jpg_path, mov_path FROM library WHERE library_code = ? and status = 0`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(code)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var l utility.Library
	if result.Next() {
		err = result.Scan(&(l.LibraryCode), &(l.LibraryName), &(l.LibraryPath), &(l.DpxPath), &(l.JpgPath), &(l.MovPath))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &l, err
}

//func QueryLibrarys(projectCode *string) (*[]string, error) {
//	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT DISTINCT material_type FROM material WHERE status = 0 AND project_code = ?`)
//	if err != nil {
//		pillarsLog.PillarsLogger.Print(err.Error())
//		return nil, err
//	}
//	defer stmt.Close()
//	result, err := stmt.Query(projectCode)
//	if err != nil {
//		pillarsLog.PillarsLogger.Print(err.Error())
//		return nil, err
//	}
//	defer result.Close()
//	var filetypes []string
//	for result.Next() {
//		var t string
//		err = result.Scan(&(t))
//		if err != nil {
//			pillarsLog.PillarsLogger.Print(err.Error())
//		}
//		filetypes = append(filetypes, t)
//	}
//	return &filetypes, err
//}
