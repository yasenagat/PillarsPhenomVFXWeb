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
		return false, err
	}

	return true, err
}

//func QueryLibraryByLibraryCode(code *string) (*utility.Library, error) {
//	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT project_code, project_name, picture, project_leader, project_type, start_datetime, end_datetime, project_detail FROM project WHERE project_code = ? and status = 0`)
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
//	var p utility.Project
//	if result.Next() {
//		err = result.Scan(&(p.ProjectCode), &(p.ProjectName), &(p.Picture), &(p.ProjectLeader), &(p.ProjectType), &(p.StartDatetime), &(p.EndDatetime), &(p.ProjectDetail))
//		if err != nil {
//			pillarsLog.PillarsLogger.Print(err.Error())
//		}
//	}
//	return &p, err
//}
