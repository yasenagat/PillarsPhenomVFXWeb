package userStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoShot(shot *utility.Shot) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO shor
		(shot_code, material_code, in_point, out_point, field_name, status) 
		VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(shot.ShotCode, shot.MaterialCode, shot.InPoint,
		shot.OutPoint, shot.FieldName, shot.Status)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func DeleteShotbyShotCode(shotCode *string) (bool, err) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE shot SET status = 1 WHERE 
		shot_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(shotCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryShotByShotCode(shotCode *string) (*utility.Shot, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT shot_code, 
		material_code, in_point, out_point, field_name, status, 
		insert_datetime, update_datetime FROM shot WHERE shot_id = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(shotCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var shot utility.Shot
	if result.Next() {
		err = result.Scan(&(shot.ShotCode), &(shot.MaterialCode),
			&(shot.InPoint), &(shot.OutPoint), &(shot.FieldName), &(shot.Status),
			&(shot.InsertDatetime), &(shot.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &shot, err
}
