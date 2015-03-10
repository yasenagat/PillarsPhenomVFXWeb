package postAction

import (
	"PillarsPhenomVFXWeb/storage/postStorage"
	"PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddRelation(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		return
	}
	var relation utility.Relation
	json.Unmarshal(data, &relation)
	name := "nRelation"
	relation.RelationCode = *utility.GenerateCode(&name)
	err = postStorage.AddRelation(&relation)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}
	/*   暂时不知需不需要
	result, err := postStorage.query(&shot.ShotCode)
	if err != nil {
		utility.OutputJson(w, 1, err.Error(), nil)
		return
	}*/
	utility.OutputJson(w, 0, "addRelation success.", nil)
}

/////  need this new struct
type moreRel struct {
	Parent string
	Shot   []string
}

func AddMoreRelation(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var slice moreRel
	var rels []utility.Relation
	json.Unmarshal(data, &slice)
	name := "nRelation"
	for i := len(slice.Shot) - 1; i >= 0; i-- {
		var rel utility.Relation

		rel.RelationCode = *utility.GenerateCode(&name)
		rel.ParentCode = slice.Parent
		rel.ChildCode = slice.Shot[i]
		rel.IsShot = true
		rel.Status = 0
		rels = append(rels, rel)
	}
	err = postStorage.AddMoreRelation(rels)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "DeleteRelation success.", nil)
}

//////////////删除关联的镜头关系！！！！！！！！！！！！！！！
func DeleteRealtion(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteRealtion(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "DeleteRelation success.", nil)
}
func DeleteRealtionByParent(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	err = postStorage.DeleteRealtionByParent(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "DeleteRelation success.", nil)
}
func QueryRelationByParentCode(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.OutputJson(w, 1, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	var code string
	json.Unmarshal(data, &code)
	result, err := postStorage.QueryRelationByParentCode(&code)
	if err != nil {
		utility.OutputJson(w, 2, "Read body failed!", nil)
		//		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}
	utility.OutputJson(w, 0, "Query success.", result)
}
