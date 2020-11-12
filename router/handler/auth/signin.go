package auth

import (
	"go-docker/helper"
	"go-docker/model"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	req := helper.ParseJsonBody(r.Body).(map[string]interface{})
	user := model.User{}
	helper.MapToStruct(req, &user)
	res, err := user.GetUserByEmail()
	if err == nil && res.Uuid == "" {
		helper.ReturnResponseAsJSON(w, nil, "Invalid Credential", 400)
		return
	} else if res.Uuid == "" && err != nil {
		helper.ErrorHandler(err, w)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req["password"].(string))); err != nil {
		helper.ReturnResponseAsJSON(w, nil, "Invalid Credential", 400)
		return
	}
	helper.ReturnResponseAsJSON(w, res, "", 200)

}
