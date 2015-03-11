package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
	"database/sql"
	"fmt"
)

func CopyEdlShot(tx *sql.Tx, n int, edls []*utility.EdlShot) ([]utility.Shot, error) {
	shots := []utility.Shot{}

	for i := n - 1; i >= 0; i-- {
		var shot utility.Shot
		shot.ShotNum = edls[i].ShotNum
		shot.ShotType = edls[i].ShotType
		shot.FromClipName = edls[i].FromClipName
		shot.SourceFile = edls[i].SourceFile
		shot.StartDateTime = edls[i].StartDateTime
		shot.EndDateTime = edls[i].EndDateTime
		shot.ShotCode = *utility.GenerateCode(&shot.ShotNum)
		//通过素材表查信息
		stmt, err := tx.Prepare("SELECT material_code, width, height, timecode_framerate FROM material WHERE status = 0 AND project_code = ? AND material_name = ?")
		if err != nil {
			return []utility.Shot{}, err
			break
		}
		result := stmt.QueryRow(shot.SourceFile)
		result.Scan(&shot.MaterialCode, &shot.Width, &shot.Height, &shot.ShotFps)
		shots = append(shots, shot)
		stmt.Close()
	}
	return shots, nil
}

//add a edl file shots start a transaction
func InsertMultipleShot(code string, edls []*utility.EdlShot) error {
	length := len(edls)
	tx, _ := mysqlUtility.DBConn.Begin()
	shots, err := CopyEdlShot(tx, length, edls)
	if err != nil {
		tx.Rollback()
		return err
	}

	for i := length - 1; i >= 0; i-- {
		stmt, err := tx.Prepare("INSERT INTO `shot`(shot_code,project_code,material_code,shot_num,start_time,end_time,clip_name,soure_file,shot_type,shot_name,shot_fps,width,height,status) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmt.Exec(shots[i].ShotCode, shots[i].ProjectCode, shots[i].MaterialCode, code, shots[i].ShotNum, shots[i].StartDateTime, shots[i].EndDateTime, shots[i].FromClipName, shots[i].SourceFile,
			shots[i].ShotType, shots[i].ShotName, shots[i].ShotFps, shots[i].Width, shots[i].Height, shots[i].Status)
		if err != nil {
			stmt.Close()
			tx.Rollback()
			return err
		}
		stmt.Close()
	}
	tx.Commit()
	return nil
}

//cover old shot   start a transaction   code is ProjectCode
func CoverMultipleShot(code string, edls []*utility.EdlShot) error {
	length := len(edls)
	tx, _ := mysqlUtility.DBConn.Begin()
	shots, err := CopyEdlShot(tx, length, edls)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt2, err2 := tx.Prepare("UPDATE `shot` SET status=1") //全部的‘删除’
	if err2 != nil {
		stmt2.Close()
		tx.Rollback()
		return err2
	}
	_, err2 = stmt2.Exec()
	if err2 != nil {
		stmt2.Close()
		tx.Rollback()
		return err2
	}
	stmt2.Close()

	for i := length - 1; i >= 0; i-- {
		stmt, err := tx.Prepare("INSERT INTO  `shot`(shot_code,project_code,material_code,shot_num,start_time,end_time,clip_name,soure_file,shot_type,shot_name,shot_fps,width,height,status) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = stmt.Exec(shots[i].ShotCode, shots[i].ProjectCode, shots[i].MaterialCode, code, shots[i].ShotNum, shots[i].StartDateTime, shots[i].EndDateTime, shots[i].FromClipName, shots[i].SourceFile,
			shots[i].ShotType, shots[i].ShotName, shots[i].ShotFps, shots[i].Width, shots[i].Height, shots[i].Status)
		if err != nil {
			stmt.Close()
			tx.Rollback()
			return err
		}
		stmt.Close()
	}
	tx.Commit()
	return nil
}

func AddSingleShot(shot *utility.Shot) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO  `shot`(shot_code,project_code,material_code,shot_num,start_time,end_time,clip_name,soure_file,shot_type,shot_name,shot_fps,width,height,status) VALUE(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(shot.ShotCode, shot.ProjectCode, shot.MaterialCode, shot.ShotNum, shot.StartDateTime, shot.EndDateTime, shot.FromClipName, shot.SourceFile,
		shot.ShotType, shot.ShotName, shot.ShotFps, shot.Width, shot.Height, shot.Status)
	if err != nil {
		stmt.Close()
		return err
	}
	return nil
}

func DeleteSingleShot(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `shot` SET status=1 WHERE shot_code=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(*code)
	if err != nil {
		stmt.Close()
		return err
	}
	return nil
}
func ModifyShotName(code *string, name *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `shot` SET shot_name=? WHERE shot_code=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, code)
	if err != nil {
		stmt.Close()
		return err
	}
	return nil
}
func QueryShotByShotCode(code *string) (*utility.Shot, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT shot_code,material_code,shot_num,start_time,end_time,clip_name,soure_file,shot_type,shot_name,shot_fps,width,height,status,insert_datetime,update_datetime FROM `shot` WHERE shot_code=?  AND status=0")
	defer stmt.Close()
	var shot utility.Shot
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&shot.ShotCode, &shot.MaterialCode, &shot.ShotNum, &shot.StartDateTime, &shot.EndDateTime, &shot.FromClipName, &shot.SourceFile,
		&shot.ShotType, &shot.ShotName, &shot.ShotFps, &shot.Width, &shot.Height, &shot.Status, &shot.InsertDatetime, &shot.UpdateDatetime)
	if err != nil {
		fmt.Println("the shot result scan appere one error")
		return nil, err
	}
	return &shot, nil
}
func QueryShotByProjectCode(code *string) ([]utility.Shot, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT shot_code,material_code,shot_num,start_time,end_time,clip_name,soure_file,shot_type,shot_name,shot_fps,width,height,status,insert_datetime,update_datetime FROM `shot` WHERE project_code= ? AND status=0")
	defer stmt.Close()
	var shots []utility.Shot
	if err != nil {
		return shots, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return shots, err
	}

	if result.Next() {
		var shot utility.Shot
		err := result.Scan(&shot.ShotCode, &shot.MaterialCode, &shot.ShotNum, &shot.StartDateTime, &shot.EndDateTime, &shot.FromClipName, &shot.SourceFile,
			&shot.ShotType, &shot.ShotName, &shot.ShotFps, &shot.Width, &shot.Height, &shot.Status, &shot.InsertDatetime, &shot.UpdateDatetime)
		if err != nil {
			fmt.Println("the shot result scan appere one error")
			return shots, err
		}
		shots = append(shots, shot)
	}
	return shots, nil
}
