package utility

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	"encoding/json"
	"fmt"
	"net/http"
)

func OutputJsonLog(w http.ResponseWriter, ret int, reason string, i interface{}, logInfo string) {
	out := &FeedbackMessage{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(b))

	if ret != 0 {
		pillarsLog.PillarsLogger.Print(logInfo)
	}
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &FeedbackMessage{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(b))
}
