package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//set fungsi ParseBody()
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//mengubah json ke dalam tipe data apapun
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
