package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

// requment struct
func AddRequment(requment *utility.Requment) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO `requment`(requment_code,shot_code,important,status,thumbnail,content) VALUE(?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(requment.RequmentCode, requment.ShotCode, requment.Important, requment.Status, requment.Image, requment.Content)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRequment(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `requment` SET status=1 WHERE requment_code=? ")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(code)
	if err != nil {
		return err
	}
	return nil
}
func QuerySingleRequment(code *string) (*utility.Requment, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT requment_code,shot_code,important,status,thumbnail,content FROM requment WHERE requment_code=?  AND status=0")
	defer stmt.Close()
	var req utility.Requment
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&req.RequmentCode, &req.ShotCode, &req.Important, &req.Status, &req.Image, &req.Content)
	if err != nil {
		//fmt.Println("the note result scan appere one error")
		return nil, err
	}
	return &req, nil
}
func QueryReqByShot(code *string) ([]utility.Requment, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT requment_code,shot_code,important,status,thumbnail,content FROM requment WHERE shot_code=?  AND status=0")
	defer stmt.Close()
	var requments []utility.Requment
	if err != nil {
		return requments, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return requments, err
	}
	if result.Next() {
		var req utility.Requment
		err = result.Scan(&req.RequmentCode, &req.ShotCode, &req.Important, &req.Status, &req.Image, &req.Content)
		if err != nil {
			return requments, err
		}
		requments = append(requments, req)

	}
	return requments, nil
}
