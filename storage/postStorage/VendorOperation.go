package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

// Note struct
func AddVendor(vendor *utility.Vendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO `vendor`(vendor_code,project_code,vendorName,content,status) VALUE(?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(vendor.VendorCode, vendor.ProjectCode, vendor.VendorName, vendor.Content, vendor.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVendor(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `vendor` SET status=1 WHERE vendor_code=? ")
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

func ModifyVendorAuth(code *string, auth *utility.VendorAuth) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `vendor` SET open_detail=?,open_demo=?,down_material=?,up_demo=?,up_product=? WHERE vendor_code=? ")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(auth.OpenDetail, auth.OpenDemo, auth.DownMaterial, auth.UpDemo, auth.UpProduct, code)
	if err != nil {
		return err
	}
	return nil
}

/* 查询单个外包商*/
func QueryVendorByVendor(code *string) (*utility.Vendor, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT vendor_code,project_code,vendorName,open_detail,open_demo,down_material,up_demo,up_product,content,status FROM vendor WHERE vendor_code=?")
	defer stmt.Close()
	var vendor utility.Vendor
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&vendor.VendorCode, &vendor.ProjectCode, &vendor.VendorName, &vendor.OpenDetail, &vendor.OpenDemo, &vendor.DownMaterial, &vendor.UpDemo, &vendor.UpProduct, &vendor.Content, &vendor.Status)
	if err != nil {
		return nil, err
	}

	return &vendor, nil
}

/* 查询外包商列表*/
func QueryVendorByPro(code *string) ([]utility.Vendor, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT vendor_code,project_code,vendorName,open_detail,open_demo,down_material,up_demo,up_product,content,status FROM vendor WHERE project_code=?")
	defer stmt.Close()
	var vendors []utility.Vendor
	if err != nil {
		return vendors, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return vendors, err
	}
	if result.Next() {
		var vendor utility.Vendor
		err = result.Scan(&vendor.VendorCode, &vendor.ProjectCode, &vendor.VendorName, &vendor.OpenDetail, &vendor.OpenDemo, &vendor.DownMaterial, &vendor.UpDemo, &vendor.UpProduct, &vendor.Content, &vendor.Status)
		if err != nil {
			return vendors, err
		}
		vendors = append(vendors, vendor)

	}
	return vendors, nil
}
