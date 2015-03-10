package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
	//"database/sql"
)

func AddRelation(relation *utility.Relation) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO  `relation`(relation_code,parent_code,child_code,list_name,isShot,content,status) VALUE(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(relation.RelationCode, relation.ParentCode, relation.ChildCode, relation.ListName, relation.IsShot, relation.Content, relation.Status)
	if err != nil {
		return err
	}
	return nil
}
func AddMoreRelation(relations []utility.Relation) error {
	tx, _ := mysqlUtility.DBConn.Begin()
	for i := len(relations) - 1; i >= 0; i-- {
		stmt, err := tx.Prepare("INSERT INTO  `relation`(relation_code,parent_code,child_code,list_name,isShot,content,status) VALUE(?,?,?,?,?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		stmt.Exec(relations[i].RelationCode, relations[i].ParentCode, relations[i].ChildCode, relations[i].ListName, relations[i].IsShot, relations[i].Content, relations[i].Status)
		stmt.Close()
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
func DeleteRealtion(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("DELETE * FROM `relation`  WHERE relation_code=?")
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
func DeleteRealtionByParent(code *string) error {
	stmt, err := mysqlUtility.DBConn.Prepare("DELETE * FROM `relation`  WHERE parent_code=?")
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
func QueryRelationByParentCode(code *string) ([]utility.Relation, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT relation_code,parent_code,child_code,list_name,isShot,content,status FROM relation WHERE parent_code=?")
	var relations []utility.Relation
	if err != nil {
		return relations, err
	}
	result, err := stmt.Query(code)
	defer result.Close()
	if err != nil {
		return relations, err
	}

	if result.Next() {
		var relation utility.Relation
		err := result.Scan(&relation.RelationCode, &relation.ParentCode, &relation.ChildCode, &relation.ListName, &relation.IsShot, &relation.Content, &relation.Status)
		if err != nil {
			println("the relation result scan appere one error")
		}
		relations = append(relations, relation)
	}
	return relations, nil
}
