package helper

import (
	"encoding/json"
	"fmt"
)

// MapToJSONByte map to json string byte
func MapToJSONByte(m interface{}) []byte {
	jsonString, _ := json.Marshal(m)
	return jsonString
}

// StringJSONToMap cconvert string json to map
func StringJSONToMap(s string) interface{} {
	var respMap interface{}
	errJSON := json.Unmarshal([]byte(s), &respMap)
	if errJSON != nil {
		fmt.Println(errJSON.Error())
		return map[string]interface{}{}
	}
	return respMap
}
