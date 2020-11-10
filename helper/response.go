package helper

import "net/http"

type Response struct {
	Message string                 `json:"message"`
	Result  map[string]interface{} `json:"result"`
}

// ReturnResponseAsJSON return response json
func ReturnResponseAsJSON(w http.ResponseWriter, resp map[string]interface{}, message string, statusCode int) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	if resp == nil {
		resp = map[string]interface{}{}
	}
	response := &Response{
		Result:  resp,
		Message: message,
	}

	w.WriteHeader(statusCode)
	return w.Write(MapToJSONByte(response))
}
