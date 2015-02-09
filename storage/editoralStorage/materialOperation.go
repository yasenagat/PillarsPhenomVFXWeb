package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertMaterial(m *utility.Material) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material (material_code, material_name, material_path, material_type, material_length, material_size, material_rate, material_start, material_end, picture, material_data, encoded_path, dpx_path, mov_path, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.MaterialCode, m.MaterialName, m.MaterialPath, m.MaterialType, m.MaterialLength, m.MaterialSize, m.MaterialRate, m.MaterialStart, m.MaterialEnd, m.Picture, m.MaterialData, m.EncodedPath, m.DpxPath, m.MovPath, m.UserCode, m.ProjectCode, m.Status)
	if err != nil {
		return false, err
	}

	return true, err
}

func DeleteMaterialByMaterialCode(materialCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material SET status = 1 WHERE material_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(materialCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}

	return true, err
}

func QueryMaterialByMaterialCode(materialCode *string) (*utility.Material, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT material_code, material_path, material_name, material_type FROM material WHERE status = 0 AND material_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(materialCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var m utility.Material
	if result.Next() {
		err = result.Scan(&(m.MaterialCode), &(m.MaterialPath), &(m.MaterialName), &(m.MaterialType), &(m.Status), &(m.InsertDatetime), &(m.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}

	return &m, err
}
