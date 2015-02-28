package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterialFolderData(d *utility.MaterialFolderData) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material_folder_data(data_code, folder_code, material_code, user_code, project_code, status, insert_datetime, update_datetime) SELECT ?, ?, ?, ?, ?, ?, NOW(), NOW() FROM DUAL WHERE NOT EXISTS(SELECT material_code FROM material_folder_data WHERE project_code = ? AND folder_code = ? AND material_code = ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.DataCode, d.FolderCode, d.MaterialCode, d.UserCode, d.ProjectCode, d.Status, d.ProjectCode, d.FolderCode, d.MaterialCode)
	if err != nil {
		return false, err
	}

	return true, err
}
