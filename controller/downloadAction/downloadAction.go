package downloadAction

import (
	s "PillarsPhenomVFXWeb/session"
	es "PillarsPhenomVFXWeb/storage/editoralStorage"
	u "PillarsPhenomVFXWeb/utility"
	"Webdemo/src/mysqlUtility"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type fileStruct struct {
	MaterialCode string
	SourceType   string
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["MaterialCode"]) + len(r.Form["SourceType"])
	if olen != 2 {
		u.OutputJsonLog(w, 1, "Error parameter format", nil, "Error parameter format")
		return
	}
	var MaterialCode = r.Form["MaterialCode"][0]
	var SourceType = r.Form["SourceType"][0]
	if len(MaterialCode) == 0 {
		u.OutputJsonLog(w, 12, "Error parameter MaterialCode", nil, "Error parameter MaterialCode")
		return
	}
	if len(SourceType) == 0 {
		u.OutputJsonLog(w, 13, "Error parameter SourceType", nil, "Error parameter SourceType")
		return
	}

	var filePath string
	var fileType string
	var libraryCode string
	var materialName string
	var materialPath string
	var materialType string
	result := mysqlUtility.DBConn.QueryRow("SELECT library_code, material_name, material_type, material_path FROM material WHERE material_code= ? AND status = 0", MaterialCode)
	err := result.Scan(&(libraryCode), &(materialName), &(materialType), &(materialPath))
	if err != nil {
		u.OutputJsonLog(w, 14, "File Not Found", nil, "File Not Found")
		return
	}
	library, err := es.QueryLibraryByLibraryCode(&libraryCode)
	if err != nil {
		u.OutputJsonLog(w, 15, "Query Material failed!", nil, "editoralStorage.QueryLibraryByLibraryCode(&libraryCode) failed!")
		return
	}
	if SourceType == "DPX" {
		filePath = library.DpxPath + materialPath + "dpx"
		fileType = "dpx"
	} else if SourceType == "JPG" {
		filePath = library.JpgPath + materialPath + "jpeg"
		fileType = "jpeg"
	} else if SourceType == "Mov" {
		filePath = library.JpgPath + materialPath + "mp4"
		fileType = "mp4"
	} else {
		filePath = library.JpgPath + materialPath + materialType
		fileType = materialType
	}

	fmt.Println("filePath---------->", filePath)
	//filePath = "/home/pillars/Videos/redone/B011_C006_06059G.RDC/B011_C006_06059G_001.R3D"
	//w.Header().Set("Content-Disposition", "attachment; filename=B011_C006_06059G_001.R3D")
	//w.Header().Set("Content-Type", "application/R3D")

	w.Header().Set("Content-Disposition", "attachment; filename="+materialName+fileType)
	w.Header().Set("Content-Type", "application/"+fileType)
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(w, file)
}

func DownloadFile1(w http.ResponseWriter, r *http.Request) {
	//flag, _ := s.GetAuthorityCode(w, r, "制片")
	//if !flag {
	//	http.Redirect(w, r, "/404.html", http.StatusFound)
	//	return
	//}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "ioutil.ReadAll(r.Body) failed!")
		return
	}

	fs := fileStruct{}
	err = json.Unmarshal(data, &fs)
	if err != nil {
		u.OutputJsonLog(w, 12, "Pramaters failed!", nil, "json.Unmarshal(data, &fs) failed!")
		return
	}

	if len(fs.MaterialCode) == 0 || len(fs.SourceType) == 0 {
		u.OutputJsonLog(w, 13, "Pramaters failed!", nil, "Pramaters failed!")
		return
	}

	var filePath string

	filePath = "/home/pillars/Videos/redone/B011_C006_06059G.RDC/B011_C006_06059G_001.R3D"
	w.Header().Set("Content-Disposition", "attachment; filename=B011_C006_06059G_001.R3D")
	w.Header().Set("Content-Type", "application/R3D")
	fmt.Println("---------->11111")
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(w, file)

	//var fileType string
	//var libraryCode string
	//var materialName string
	//var materialPath string
	//var materialType string
	//result := mysqlUtility.DBConn.QueryRow("SELECT library_code, material_name, material_type, material_path FROM material WHERE material_code= ? AND status = 0", fs.MaterialCode)
	//result.Scan(&(libraryCode), &(materialName), &(materialType), &(materialPath))
	//library, err := es.QueryLibraryByLibraryCode(&libraryCode)
	//if err != nil {
	//	u.OutputJsonLog(w, 14, "Query Material failed!", nil, "editoralStorage.QueryLibraryByLibraryCode(&libraryCode) failed!")
	//	return
	//}
	//if fs.SourceType == "DPX" {
	//	filePath = library.DpxPath + materialPath + "dpx"
	//	fileType = "dpx"
	//} else if fs.SourceType == "JPG" {
	//	filePath = library.JpgPath + materialPath + "jpeg"
	//	fileType = "jpeg"
	//} else if fs.SourceType == "Mov" {
	//	filePath = library.JpgPath + materialPath + "mp4"
	//	fileType = "mp4"
	//} else {
	//	filePath = library.JpgPath + materialPath + materialType
	//	fileType = materialType
	//}
	//w.Header().Set("Content-Disposition", "attachment; filename="+materialName+fileType)
	//w.Header().Set("Content-Type", "application/"+fileType)
	//file, _ := os.Open(filePath)
	//defer file.Close()
	//io.Copy(w, file)
}
