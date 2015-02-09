package editoralAction

import (
	s "PillarsPhenomVFXWeb/session"
	es "PillarsPhenomVFXWeb/storage/editoralStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddLibrary(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["LibraryName"]) + len(r.Form["LibraryPath"]) + len(r.Form["EncodedPath"]) + len(r.Form["DpxPath"]) + len(r.Form["MovPath"])
	if olen != 6 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}

	if len(r.Form["LibraryName"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter LibraryName", nil)
		return
	}

	if len(r.Form["LibraryPath"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter LibraryPath", nil)
		return
	}

	temp := "insert"
	code := u.GenerateCode(&temp)
	library := u.Library{
		LibraryCode: *code,
		LibraryName: r.Form["LibraryName"][0],
		LibraryPath: r.Form["LibraryPath"][0],
		EncodedPath: r.Form["EncodedPath"][0],
		DpxPath:     r.Form["DpxPath"][0],
		MovPath:     r.Form["MovPath"][0],
		UserCode:    userCode,
		ProjectCode: r.Form["ProjectCode"][0],
		Status:      0,
	}
	result, _ := es.InsertLibrary(&library)
	if result == false {
		u.OutputJson(w, 15, "Insert into library failed!", nil)
		return
	}

	err, materials := LoadMaterials(library.LibraryPath)
	if err != nil {
		u.OutputJson(w, 16, "Load materials failed!", nil)
		return
	}
	for n, m := range materials {
		// TODO 调用C++，传入素材路径，返回素材的信息
		fmt.Println(n, "-------->", m.MaterialPath)

		// -----------temp start-----------
		m.MaterialCode = *u.GenerateCode(&temp)
		m.DpxPath = "DpxPath测试数据"
		m.MaterialLength = "len测试数据"
		m.MaterialSize = "size测试数据"
		m.MaterialRate = "rate测试数据"
		m.MaterialStart = "start测试数据"
		m.MaterialEnd = "end测试数据"
		m.Picture = "Picture测试数据"
		m.MaterialData = "MaterialData测试数据"
		m.EncodedPath = "EncodedPath测试数据"
		m.DpxPath = "DpxPath测试数据"
		m.MovPath = "MovPath测试数据"
		m.UserCode = library.UserCode
		m.ProjectCode = library.ProjectCode
		m.Status = 0
		// -----------temp end-----------

		result, _ := es.InsertMaterial(m)
		if result == false {
			u.OutputJson(w, 17, "Save material failed!", nil)
			return
		}
	}

	rs, _ := json.Marshal(library)
	u.OutputJson(w, 0, "Add library succeed!", string(rs))
}

func GetLibraryFileList(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["LibraryCode"]) + len(r.Form["Start"]) + len(r.Form["End"])
	if olen != 4 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}

	if len(r.Form["LibraryCode"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter LibraryCode", nil)
		return
	}

	start, err := strconv.ParseInt(r.Form["Start"][0], 10, 0)
	if err != nil {
		u.OutputJson(w, 14, "Error parameter Start", nil)
		return
	}

	end, err := strconv.ParseInt(r.Form["End"][0], 10, 0)
	if err != nil {
		u.OutputJson(w, 15, "Error parameter End", nil)
		return
	}

	// TODO -----------------------------------------------------未完待续
	projectList, err := es.QueryMaterialList(r.Form["LibraryCode"][0], start, end)
	if err != nil {
		u.OutputJson(w, 16, "Query MaterialList failed!", nil)
		return
	}

	rs, _ := json.Marshal(projectList)
	u.OutputJson(w, 0, "Load project succeed!", string(rs))
}
