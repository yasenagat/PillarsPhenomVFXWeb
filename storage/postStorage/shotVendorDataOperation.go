package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertShotVendorData(d *utility.ShotVendorData) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_vendor_data(data_code, vendor_code, vendor_user, shot_code, project_code, user_code, status, insert_datetime, update_datetime) SELECT ?, ?, ?, ?, ?, ?, ?, NOW(), NOW() FROM DUAL WHERE NOT EXISTS(SELECT shot_code FROM shot_vendor_data WHERE status = 0 AND project_code = ? AND vendor_code = ? AND shot_code = ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.DataCode, d.VendorCode, d.VendorUser, d.ShotCode, d.ProjectCode, d.UserCode, d.Status, d.ProjectCode, d.VendorCode, d.ShotCode)
	if err != nil {
		return err
	}

	return nil
}

func DeleteShotVendorData(userCode string, projectCode string, vendorCode string, shotCodes string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor_data SET user_code = ?, status = 1, update_datetime = NOW() WHERE status = 0 AND project_code = ? AND vendor_code = ? AND shot_code in (" + shotCodes + ")")
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userCode, projectCode, vendorCode)
	if err != nil {
		return false, err
	}

	return true, err
}
