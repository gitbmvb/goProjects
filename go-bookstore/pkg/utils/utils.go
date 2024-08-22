package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r * http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(r.body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
	
}