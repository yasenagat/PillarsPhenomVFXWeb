package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddThumbnail(thumbnail *utility.Thumbnail) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO  `thumbnail`(thumbnail_code,shot_code,thumbnail_image,status) VALUE(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(thumbnail.ThumbnailCode, thumbnail.ShotCode, thumbnail.ThumbnailImage, 0)
	if err != nil {
		return err
	}
	return nil
}

func DeleteThumbnail(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `thumbnail` SET status=1 WHERE thumbnail_code=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(code)
	if err != nil {
		return err
	}
	return nil
}
func QuerySingleImage(code *string) (*utility.Thumbnail, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT thumbnail_code,shot_code,thumbnail_image FROM thumbnail WHERE thumbnail_code=? AND status=0")
	var thumbnail utility.Thumbnail
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&thumbnail.ThumbnailCode, &thumbnail.ShotCode, &thumbnail.ThumbnailImage)
	if err != nil {
		println("the thumbnail result scan appere one error")
		return nil, err
	}
	return &thumbnail, nil
}
func QueryThumbnailByShotCode(code *string) ([]utility.Thumbnail, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT thumbnail_code,shot_code,thumbnail_image FROM thumbnail WHERE shot_code=? AND status=0")
	var thumbnails []utility.Thumbnail
	if err != nil {
		return thumbnails, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return thumbnails, err
	}

	if result.Next() {
		var thumbnail utility.Thumbnail
		err := result.Scan(&thumbnail.ThumbnailCode, &thumbnail.ShotCode, &thumbnail.ThumbnailImage)
		if err != nil {
			println("the thumbnail result scan appere one error")
			return thumbnails, err
		}
		thumbnails = append(thumbnails, thumbnail)
	}
	return thumbnails, nil
}
