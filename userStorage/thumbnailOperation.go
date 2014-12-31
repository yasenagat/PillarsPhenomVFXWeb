package userStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoThumbnail(thumbnail *utility.Thumbnail) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO thumbnail
		(thumbnail_code, shot_code, thumbnail_image, status) VALUES(?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(thumbnail.ThumbnailCode, thumbnail.ShotCode,
		thumbnail.ThumbnailImage, thumbnail.Status)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func DeleteThumbnailByThumbnailCode(thumbnailCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE thumbnail SET status = 1 
		WHERE thumbnail_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(thumbnailCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryThumbnailByThumbnailCode(thumbnailCode *string) (*utility.Thumbnail, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT thumbnail_code, 
		shot_code, thumbnail_image, status, insert_datetime, update_datetime 
		FROM thumbnail WHERE shot_id = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(thumbnailCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var thumbnail utility.Thumbnail
	if result.Next() {
		err = result.Scan(&(thumbnail.thumbnailCode), &(thumbnail.ShotCode),
			&(thumbnail.ThumbnailImage), &(thumbnail.Status),
			&(shot.InsertDatetime), &(shot.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &thumbnail, err
}
