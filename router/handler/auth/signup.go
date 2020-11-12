package auth

import (
	"go-docker/helper"
	"go-docker/model"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	req := helper.ParseJsonBody(r.Body).(map[string]interface{})

	uuidVal, _ := uuid.NewUUID()

	user := &model.User{
		Uuid:      uuidVal.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	helper.MapToStruct(req, &user)

	err := helper.ValidateStruct(user)
	if err != nil {
		helper.ErrorHandler(err, w)
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.ErrorHandler(err, w)
		return
	}
	user.Password = string(password)
	uuidValMerchant, _ := uuid.NewUUID()
	merchant := &model.Merchant{
		Uuid:      uuidValMerchant.String(),
		UserUuid:  user.Uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	helper.MapToStruct(req, &merchant)
	user.CreateUser(*merchant)
	helper.ReturnResponseAsJSON(w, nil, "Success Create User", 200)
}
