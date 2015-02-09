package pillarsLog

import (
	"PillarsPhenomVFXWeb/utility"
	"fmt"
	"log"
	"os"
)

var PillarsLogger *log.Logger
var outFile *os.File

func init() {
	if PillarsLogger == nil {
		propertyMap := utility.ReadProperty("../log.properties")
		logFileName := propertyMap["LogFile"]
		fmt.Println(logFileName)
		var err error
		outFile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			panic(err.Error())
		}
		PillarsLogger = log.New(outFile, "", log.Ldate|log.Ltime|log.Llongfile)
	}
}

func CloseLogFile() {
	outFile.Close()
}
