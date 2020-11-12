package helper

import (
	"encoding/json"
	"fmt"
	"io"
)

// MapToJSONByte map to json string byte
func MapToJSONByte(m interface{}) []byte {
	jsonString, _ := json.Marshal(m)
	return jsonString
}

// MapToJSONByte map to json string byte
func MapToStruct(m interface{}, s interface{}) {
	json.Unmarshal(MapToJSONByte(m), &s)
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

func ParseJsonBody(body io.ReadCloser) interface{} {
	var result interface{}
	json.NewDecoder(body).Decode(&result)
	return result
}
