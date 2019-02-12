package stringobj

import (
	"encoding/json"
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	logger = log.NewAppLogger("pkg/utils/stringobj")
)

// Map2list is for exchanging map to list
func Map2list(m map[string]string) string {
	s := ""
	for k, v := range m {
		s = s + fmt.Sprintf("%s:%s ", k, v)
	}
	return s
}

// String2Json is for exchanging string to json
func String2Json(data []byte) []byte {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	data, err = json.MarshalIndent(v, "", " ")
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	return data
}
