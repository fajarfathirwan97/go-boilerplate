package helper

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(err error, w http.ResponseWriter) {
	if reflect.TypeOf(err).Name() != "ValidationErrors" {
		logrus.Errorln(err)
		ReturnResponseAsJSON(w, nil, "Oops something went wrong", 500)
		return
	}
	var eMsg []map[string]interface{}
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "email":
			eMsg = append(eMsg, map[string]interface{}{
				err.Field(): fmt.Sprintf("%v must be an %v", err.Field(), err.Tag()),
			})
		default:
			eMsg = append(eMsg, map[string]interface{}{
				err.Field(): fmt.Sprintf("%v is %v", err.Field(), err.Tag()),
			})
		}
	}
	ReturnResponseAsJSON(w, eMsg, "Error Validation", 422)
	return
}
