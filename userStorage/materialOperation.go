package userStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoMaterial(material *utility.Material) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material
		(material_code, material_path, material_name, material_type, 
		material_encoded_path, status) VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(material.MaterialCode, material.MaterialPath,
		material.MaterialName, material.MaterialType,
		material.MaterialEncodedPath, material.Status)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func DeleteMaterialByMaterialCode(materialCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material 
		SET status = 1 WHERE material_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(materialCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryMaterialByMaterialCode(materialCode *string) (*utility.Material, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT material_code, material_path,  
		material_name, material_type, material_encoded_path, status, insert_datetime, 
		update_datetime FROM material WHERE material_code = ?`)
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
	var material utility.Material
	if result.Next() {
		err = result.Scan(&(material.MaterialCode), &(material.MaterialPath),
			&(material.MaterialName), &(material.MaterialType),
			&(material.MaterialEncodedPath), &(material.Status),
			&(material.InsertDatetime), &(material.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &material, err
}
