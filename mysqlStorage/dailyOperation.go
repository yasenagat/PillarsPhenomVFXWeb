package mysqlStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoDaily(daily *utility.Daily) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO daily(daily_code, 
		shot_code, award_code, target_path, status) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(daily.DailyCode, daily.ShotCode, daily.AwardCode,
		daily.TargetPath, daily.Status)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func DeleteDailyByDailyCode(dailyCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE daily SET status = 1 
		WHERE daily_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dailyCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryDailyByDailyCode(dailyCode *string) (*utility.Daily, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT daily_code, shot_code, 
		award_code, target_path, status, insert_datetime, update_datetime 
		FROM daily WHERE daily_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	var daily utility.Daily
	result, err := stmt.Query(dailyCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&(daily.DailyCode), &(daily.ShotCode),
			&(daily.AwardCode), &(daily.TargetPath), &(daily.Status),
			&(daily.InsertDatetime), &(daily.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &daily, err
}
