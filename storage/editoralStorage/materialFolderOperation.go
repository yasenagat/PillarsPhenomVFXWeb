package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterialFolder(g *utility.MaterialFolder) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material_folder (folder_code, folder_name, father_code, leaf_flag, folder_Detail, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.FolderCode, g.FolderName, g.FatherCode, g.LeafFlag, g.FolderDetail, g.UserCode, g.ProjectCode, g.Status)
	if err != nil {
		return false, err
	}

	return true, err
}

func DeleteMaterialFolder(folderCode string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material_folder SET status = 1 WHERE status = 0 AND (folder_code = ? OR father_code = ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(folderCode, folderCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	// 删除素材组成功后，继续删除数据表的数据
	stmt, err = mysqlUtility.DBConn.Prepare(`UPDATE material_folder_data SET status = 1 WHERE status = 0 AND folder_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(folderCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}

	return true, err
}

func UpdateMaterialFolder(g *utility.MaterialFolder) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material_folder SET folder_name = ?, father_code = ?, folder_detail = ?, user_code = ?, update_datetime = now() WHERE folder_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.FolderName, g.FatherCode, g.FolderDetail, g.UserCode, g.FolderCode)
	if err != nil {
		return false, err
	}

	return true, err
}
