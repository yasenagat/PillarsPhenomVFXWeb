package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddShotDemo(sv *utility.ShotVersion) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_version(version_code, shot_code, vendor_user, version_num, picture, demo, demo_detail, product, product_detail, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.VersionCode, sv.ShotCode, sv.VendorUser, sv.VersionNum, sv.Picture, sv.Demo, sv.DemoDetail, sv.Product, sv.ProductDetail, sv.Status)
	if err != nil {
		return err
	}
	return nil
}

func GetShotVersionNum(sv *utility.ShotVersion) (*int, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT MAX(version_num) FROM shot_version WHERE status = 0 AND shot_code = ? AND vendor_user = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var n int
	result := stmt.QueryRow(sv.ShotCode, sv.VendorUser)
	result.Scan(&n)
	return &n, nil
}

func QueryShotVersion(shotCode *string) (*[]utility.ShotVersion, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT version_code, version_num, picture, demo_datail FROM shot_version WHERE status = 0 AND shot_code = ? ORDER BY version_num DESC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(shotCode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var versions []utility.ShotVersion
	for result.Next() {
		var version utility.ShotVersion
		err = result.Scan(&version.VersionCode, &version.VersionNum, &version.Picture, &version.DemoDetail)
		if err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	return &versions, nil
}

//func DeleteShotVersion(code *string) error {
//	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_version SET status = 1 WHERE status = 0 AND version_code = ?")
//	defer stmt.Close()
//	if err != nil {
//		return err
//	}
//	_, err = stmt.Exec(code)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func QuerySingleShotVersion(code *string) (*utility.ShotVersion, error) {
//	stmt, err := mysqlUtility.DBConn.Prepare("SELECT version_code,shot_code,vendor_code,thumbnail,demo,product,status FROM shot_version WHERE version_code=?  AND status=0")
//	var shot_version utility.ShotVersion
//	if err != nil {
//		return nil, err
//	}
//	result := stmt.QueryRow(code)
//	err = result.Scan(&version.VersionCode, &version.ShotCode, &version.VendorCode, &version.Image, &version.Demo, &version.Product, &version.Status)
//	if err != nil {
//		return nil, err
//	}
//	return &version, nil
//}
