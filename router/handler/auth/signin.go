package auth

import (
	"go-docker/helper"
	"go-docker/model"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	req := helper.ParseJsonBody(r.Body).(map[string]interface{})
	if db, err := helper.GetDBClient(); err != nil {
		helper.ErrorHandler(err, w)
		return
	} else {
		defer db.Close()
		user := model.User{}
		helper.MapToStruct(req, &user)
		err := db.Get(&user, "SELECT * FROM users WHERE email=$1", user.Email)
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req["password"].(string)))
		if err != nil {
			helper.ReturnResponseAsJSON(w, nil, "Invalid Credentials", 400)
			return
		}
		user.Password = ""
		helper.ReturnResponseAsJSON(w, user, "", 200)
	}
}
