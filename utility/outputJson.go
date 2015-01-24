package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &FeedbackMessage{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	fmt.Fprintf(w, string(b))
}
