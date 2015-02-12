package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterialFolder(g *utility.MaterialFolder) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material_folder (folder_code, folder_name, father_code, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.FolderCode, g.FolderName, g.FatherCode, g.UserCode, g.ProjectCode, g.Status)
	if err != nil {
		return false, err
	}

	return true, err
}
