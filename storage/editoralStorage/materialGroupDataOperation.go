package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterialGroupData(d *utility.MaterialGroupData) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material_group_data (data_code, group_code, material_code, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.DataCode, d.GroupCode, d.MaterialCode, d.UserCode, d.ProjectCode, d.Status)
	if err != nil {
		return false, err
	}

	return true, err
}
