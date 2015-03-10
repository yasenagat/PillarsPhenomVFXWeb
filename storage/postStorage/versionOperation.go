package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

// version struct
func AddVersion(version *utility.Version) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO `version`(version_code,shot_code,vendor_code,thumbnail,demo,product,status) VALUE(?,?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(version.VersionCode, version.ShotCode, version.VendorCode, version.Image, version.Demo, version.Product, version.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVersion(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `version` SET status=1 WHERE version_code=? ")
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
func QuerySingleVersion(code *string) (*utility.Version, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT version_code,shot_code,vendor_code,thumbnail,demo,product,status FROM version WHERE version_code=?  AND status=0")
	var version utility.Version
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&version.VersionCode, &version.ShotCode, &version.VendorCode, &version.Image, &version.Demo, &version.Product, &version.Status)
	if err != nil {
		println("the version result scan appere one error")
		return nil, err
	}
	return &version, nil
}
func QueryVersionByShot(code *string) ([]utility.Version, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT version_code,shot_code,vendor_code,thumbnail,demo,product,status FROM version WHERE shot_code=?  AND status=0")
	defer stmt.Close()
	var versions []utility.Version
	if err != nil {
		return versions, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return versions, err
	}
	if result.Next() {
		var version utility.Version
		err = result.Scan(&version.VersionCode, &version.ShotCode, &version.VendorCode, &version.Image, &version.Demo, &version.Product, &version.Status)
		if err != nil {
			return versions, err
		}
		versions = append(versions, version)

	}
	return versions, nil
}
