package awesome_structs

import (
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func StringMap(v interface{}) (sm map[string]string, err error) {
	m, err := json.Marshal(v)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = json.Unmarshal(m, &sm)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
