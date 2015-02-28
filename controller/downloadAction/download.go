package downloadAction

import (
	s "PillarsPhenomVFXWeb/session"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	// SELECT file地址

	//f, err := os.OpenFile("./file.exe", os.O_RDWR, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//stat, err := f.Stat() //获取文件状态
	//if err != nil {
	//	panic(err)
	//}
	//f.Seek(stat.Size, 0) //把文件指针指到文件末。
	//url := "http://dl.google.com/chrome/install/696.57/chrome_installer.exe"
	//var req http.Request
	//req.Method = "GET"
	//req.UserAgent = "Download"
	//req.Close = true
	//req.URL, err = http.ParseURL(url)
	//if err != nil {
	//	panic(err)
	//}
	//header := http.Header{}
	//header.Set("Range", "bytes="+strconv.Itoa64(stat.Size)+"-")
	//req.Header = header
	//resp, err := http.DefaultClient.Do(&req)
	//if err != nil {
	//	panic(err)
	//}
	//written, err := io.Copy(f, resp.Body)
	//if err != nil {
	//	panic(err)
	//}

}
