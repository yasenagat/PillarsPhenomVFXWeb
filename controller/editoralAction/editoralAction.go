package editoralAction

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	r3d "PillarsPhenomVFXWeb/r3dOperation"
	s "PillarsPhenomVFXWeb/session"
	es "PillarsPhenomVFXWeb/storage/editoralStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func GetLibrarys(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}
	code := r.Form["ProjectCode"][0]
	filetypes, err := es.QueryLibrarys(&code)
	if err != nil {
		u.OutputJson(w, 13, "Query Filetypes failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Query Filetypes succeed!", filetypes)
}

func AddLibrary(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["LibraryName"]) + len(r.Form["LibraryPath"]) + len(r.Form["DpxPath"]) + len(r.Form["JpgPath"]) + len(r.Form["MovPath"])
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

	_, err := os.Stat(r.Form["LibraryPath"][0])
	if err != nil {
		u.OutputJson(w, 15, "LibraryPath failed", nil)
		return
	}

	temp := "insert"
	code := u.GenerateCode(&temp)
	library := u.Library{
		LibraryCode: *code,
		LibraryName: r.Form["LibraryName"][0],
		LibraryPath: r.Form["LibraryPath"][0],
		DpxPath:     r.Form["DpxPath"][0],
		JpgPath:     r.Form["JpgPath"][0],
		MovPath:     r.Form["MovPath"][0],
		UserCode:    userCode,
		ProjectCode: r.Form["ProjectCode"][0],
		Status:      0,
	}
	result, err := es.InsertLibrary(&library)
	if result == false {
		fmt.Println(err.Error())
		u.OutputJson(w, 16, "Insert into library failed!", nil)
		return
	}

	err, materials := LoadMaterials(library.LibraryPath)
	if err != nil {
		u.OutputJson(w, 17, "Load materials failed!", nil)
		return
	}
	for _, m := range materials {
		filePath := library.LibraryPath + m.MaterialPath + m.MaterialType
		//fmt.Println(filePath)
		// TODO 调用C++，传入素材路径，返回素材的信息(图片尚未实现)
		clip := r3d.ClipInit(filePath)
		m.LibraryCode = library.LibraryCode
		m.MaterialCode = *u.GenerateCode(&temp)
		m.VideoTrackCount = r3d.ClipVideoTrackCount(clip)
		m.Width = r3d.ClipWidth(clip)
		m.Height = r3d.ClipHeight(clip)
		m.VideoAudioFramerate = r3d.ClipVideoAudioFramerate(clip)
		m.StartAbsoluteTimecode = r3d.ClipStartAbsoluteTimecode(clip)
		m.EndAbsoluteTimecode = r3d.ClipEndAbsoluteTimecode(clip)
		m.MetaData = r3d.ClipMetaData(clip)
		m.Picture = "Picture测试数据"
		m.UserCode = library.UserCode
		m.ProjectCode = library.ProjectCode
		m.Status = 0

		result, _ := es.InsertMaterial(m)
		if result == false {
			u.OutputJson(w, 18, "Save material failed!", nil)
			return
		}
	}

	u.OutputJson(w, 0, "Add library succeed!", library)
}

func GetLibraryFileList(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["Flag"]) + len(r.Form["ProjectCode"]) + len(r.Form["LibraryCode"]) + len(r.Form["Start"]) + len(r.Form["End"])
	if olen != 5 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Flag"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Flag", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter ProjectCode", nil)
		return
	}

	if len(r.Form["LibraryCode"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter LibraryCode", nil)
		return
	}

	start, err := strconv.ParseInt(r.Form["Start"][0], 10, 0)
	if err != nil {
		u.OutputJson(w, 15, "Error parameter Start", nil)
		return
	}

	end, err := strconv.ParseInt(r.Form["End"][0], 10, 0)
	if err != nil {
		u.OutputJson(w, 16, "Error parameter End", nil)
		return
	}

	var materials interface{}
	if r.Form["Flag"][0] == "1" {
		materials, err = es.QueryMaterialsByLibraryCode(r.Form["LibraryCode"][0], start, end)
	} else if r.Form["Flag"][0] == "2" {
		materials, err = es.QueryMaterialsByType(r.Form["ProjectCode"][0], r.Form["LibraryCode"][0], start, end)
	}

	if err != nil {
		u.OutputJson(w, 17, "Query Materials failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Query Materials succeed!", materials)
}

func FindMaterials(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["Args"])
	if olen != 2 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}

	if len(r.Form["Args"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter Args", nil)
		return
	}

	materials, err := es.FindMaterials(r.Form["ProjectCode"][0], r.Form["Args"][0])
	if err != nil {
		u.OutputJson(w, 13, "Find Materials failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Find Materials succeed!", materials)
}

func GetMaterialInfo(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["MaterialCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	if len(r.Form["MaterialCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter MaterialCode", nil)
		return
	}
	code := r.Form["MaterialCode"][0]
	material, err := es.QueryMaterialByMaterialCode(&code)
	if err != nil {
		u.OutputJson(w, 13, "Query Material failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Load project succeed!", material)
}

func GetFiletypes(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}
	code := r.Form["ProjectCode"][0]
	filetypes, err := es.QueryFiletypes(&code)
	if err != nil {
		u.OutputJson(w, 13, "Query Filetypes failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Query Filetypes succeed!", filetypes)
}

func chectString(w http.ResponseWriter, r *http.Request, num int, args []string) bool {
	for _, str := range args {
		if len(r.Form[str][0]) == 0 {
			u.OutputJson(w, num, "Error parameter "+str, nil)
			return false
		}
		num += 1
	}
	return true
}

func AddFolder(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["FolderName"]) + len(r.Form["FatherCode"]) + len(r.Form["FolderDetail"])
	if olen != 4 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	var args = []string{"ProjectCode", "FolderName", "FatherCode", "FolderDetail"}
	if !chectString(w, r, 12, args) {
		return
	}

	temp := "insert"
	code := u.GenerateCode(&temp)
	mf := u.MaterialFolder{
		FolderCode:   *code,
		FolderName:   r.Form["FolderName"][0],
		FatherCode:   r.Form["FatherCode"][0],
		LeafFlag:     "0",
		FolderDetail: r.Form["FolderDetail"][0],
		UserCode:     userCode,
		ProjectCode:  r.Form["ProjectCode"][0],
		Status:       0,
	}
	result, _ := es.InsertMaterialFolder(&mf)
	if result == false {
		u.OutputJson(w, 16, "Insert into material_folder failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Add material_folder succeed!", mf)
}

func DeleteFolder(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["FolderCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["FolderCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter FolderCode", nil)
		return
	}

	result, _ := es.DeleteMaterialFolder(r.Form["FolderCode"][0])
	if result == false {
		u.OutputJson(w, 13, "Delete material_folder failed!", nil)
		return
		u.OutputJson(w, 16, "Insert into material_folder failed!", nil)
	}

	u.OutputJson(w, 0, "Delete material_folder succeed!", nil)
}

func UpdateFolder(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["FolderCode"]) + len(r.Form["FolderName"]) + len(r.Form["FatherCode"]) + len(r.Form["FolderDetail"])
	if olen != 4 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	var args = []string{"FolderCode", "FolderName", "FatherCode", "FolderDetail"}
	if !chectString(w, r, 12, args) {
		return
	}

	mf := u.MaterialFolder{
		FolderCode:   r.Form["FolderCode"][0],
		FolderName:   r.Form["FolderName"][0],
		FatherCode:   r.Form["FatherCode"][0],
		FolderDetail: r.Form["FolderDetail"][0],
		UserCode:     userCode,
	}
	result, _ := es.UpdateMaterialFolder(&mf)
	if result == false {
		u.OutputJson(w, 16, "Update material_folder failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Update material_folder succeed!", mf)
}

// 请求传入JSON格式数据，使用结构体解析
type addFiles struct {
	ProjectCode   string
	FolderCode    string
	MaterialCodes []string
}

func AddFolderFiles(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJson(w, 1, "Read body failed!", nil)
		pillarsLog.PillarsLogger.Print("ioutil.ReadAll(r.Body) failed!")
		return
	}

	folderFiles := addFiles{}
	err = json.Unmarshal(data, &folderFiles)
	if err != nil {
		u.OutputJson(w, 12, "Pramaters failed!", nil)
		pillarsLog.PillarsLogger.Print("json.Unmarshal(data, &folderFiles) failed!")
		return
	}

	if len(folderFiles.ProjectCode) == 0 || len(folderFiles.FolderCode) == 0 || len(folderFiles.MaterialCodes) == 0 {
		u.OutputJson(w, 13, "Pramaters failed!", nil)
		pillarsLog.PillarsLogger.Print("Pramaters failed!")
		return
	}
	temp := "insert"
	for i, value := range folderFiles.MaterialCodes {
		mfd := u.MaterialFolderData{
			DataCode:     *u.GenerateCode(&temp),
			FolderCode:   folderFiles.FolderCode,
			MaterialCode: value,
			UserCode:     userCode,
			ProjectCode:  folderFiles.ProjectCode,
			Status:       0,
		}
		result, _ := es.InsertMaterialFolderData(&mfd)
		if result == false {
			b, _ := strconv.Atoi("1" + string(4+i))
			fmt.Println("Insert failed-------->", b)
			u.OutputJson(w, b, "Insert into material_folder_data failed!", nil)
			return
		}
	}

	u.OutputJson(w, 0, "Insert into material_folder_data succeed!", nil)
}
