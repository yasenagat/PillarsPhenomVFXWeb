package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterialGroup(g *utility.MaterialGroup) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material_group (group_code, group_name, father_code, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.GroupCode, g.GroupName, g.FatherCode, g.UserCode, g.ProjectCode, g.Status)
	if err != nil {
		return false, err
	}

	return true, err
}
