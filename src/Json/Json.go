package Json

import (
	"bytes"
	"encoding/json"
)

func ParseArray(jsonStr string) []map[string]interface{} {
	dict := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(jsonStr), &dict); err == nil {
		return dict
	}
	panic("error")
}

func Parse(jsonStr string) map[string]interface{} {
	dict := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonStr), &dict); err == nil {
		return dict
	}
	panic("error")
}
func Stringify(obj interface{}, indent bool) string {
	if jsonBytes, err := json.Marshal(obj); err == nil {
		if indent {
			var out bytes.Buffer
			json.Indent(&out, jsonBytes, "", "\t")
			return out.String()
		}
		return string(jsonBytes)
	}
	panic("error")
}
