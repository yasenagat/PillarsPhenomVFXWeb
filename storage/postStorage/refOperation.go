package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddReference(ref *utility.Reference) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO `reference`(reference_code,shot_code,reference_type,mat_name,mat_type,mat_content,mat_url,status) VALUE(?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ref.ReferenceCode, ref.ShotCode, ref.ReferenceType, ref.MatName, ref.MatType, ref.MatContent, ref.MatUrl, ref.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteReference(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `reference` SET status=1 WHERE reference_code=? ")
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
func ModifyReference(ref *utility.Reference) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE `reference` SET mat_name=?,mat_type=?,mat_content=?,mat_url=? WHERE reference_code=? AND status=0")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ref.MatName, ref.MatType, ref.MatContent, ref.MatUrl)
	if err != nil {
		return err
	}
	return nil
}
func QuerySingleReference(code *string) (*utility.Reference, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT reference_code,shot_code,reference_type,mat_name,mat_type,mat_content,mat_url,status FROM reference WHERE reference_code=?  AND status=0")
	defer stmt.Close()
	var ref utility.Reference
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&ref.ReferenceCode, &ref.ShotCode, &ref.ReferenceType, &ref.MatName, &ref.MatType, &ref.MatContent, &ref.MatUrl, &ref.Status)
	if err != nil {
		//fmt.Println("the note result scan appere one error")
		return nil, err
	}
	return &ref, nil
}
func QueryReferenceByShot(code *string) ([]utility.Reference, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT reference_code,shot_code,reference_type,mat_name,mat_type,mat_content,mat_url,status FROM reference WHERE shot_code=?  AND status=0")
	defer stmt.Close()
	var refs []utility.Reference
	if err != nil {
		return refs, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return refs, err
	}
	if result.Next() {
		var ref utility.Reference
		err = result.Scan(&ref.ReferenceCode, &ref.ShotCode, &ref.ReferenceType, &ref.MatName, &ref.MatType, &ref.MatContent, &ref.MatUrl, &ref.Status)
		if err != nil {
			return refs, err
		}
		refs = append(refs, ref)

	}
	return refs, nil
}
func QueryRefDetail(code *string) (*utility.ReferenceDetail, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT material_path,material_type,picture FROM material WHERE material_code=?  AND status=0")
	defer stmt.Close()
	var ref utility.ReferenceDetail
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(code)
	err = result.Scan(&ref.MatSource, &ref.MatDpx, &ref.MatJPG)
	if err != nil {
		//fmt.Println("the note result scan appere one error")
		return nil, err
	}
	return &ref, nil
}
