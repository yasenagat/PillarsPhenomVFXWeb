package mysqlStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoAward(award *utility.Award) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO award 
		(award_code, shot_code, recieve_user_code, require, target_path_base, 
		addition_path, status) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(award.AwardCode, award.ShotCode, award.RecieveUserCode,
		award.Require, award.TargetPathBase, award.AdditionPath, award.Status)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func DeleteAwardByAwardCode(awardCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE award SET status = 1 
		WHERE award_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(awardCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryAwardByAwardCode(awardCode *string) (*utility.Award, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT award_code, shot_code, 
		recieve_user_code, require, target_path_base, addition_path, status, 
		insert_datetime, update_datetime FROM award WHERE award_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(awardCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var award utility.Award
	if result.Next() {
		err = result.Scan(&(award.AwardCode), &(award.ShotCode),
			&(award.RecieveUserCode), &(award.Require),
			&(award.TargetPathBase), &(award.AdditionPath), &(award.Status),
			&(award.InsertDatetime), &(award.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &award, err
}
