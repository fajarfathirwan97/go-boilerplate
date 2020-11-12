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

	db, err := helper.GetDBClient()
	if err != nil {
		helper.ErrorHandler(err, w)
		return
	}
	defer db.Close()
	uuidVal, _ := uuid.NewUUID()

	user := &model.User{
		Uuid:      uuidVal.String(),
		CreatedAt: time.Now(),
	}
	helper.MapToStruct(req, &user)

	if err = helper.ValidateStruct(user); err != nil {
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
	}
	helper.MapToStruct(req, &merchant)

	tx := db.MustBegin()
	if _, err = tx.NamedExec(`
		INSERT INTO users (
			uuid,
			email,
			password,
			created_at,
			updated_at
		) VALUES (
			:uuid,
			:email,
			:password,
			:created_at,
			:updated_at
		)`, user); err != nil {
		tx.Rollback()
		helper.ErrorHandler(err, w)
		return
	}
	if _, err = tx.NamedExec(`
		INSERT INTO merchants (
			uuid,
			user_uuid,
			merchant_name,
			merchant_address,
			created_at,
			updated_at
		) VALUES (
			:uuid,
			:user_uuid,
			:merchant_name,
			:merchant_address,
			:created_at,
			:updated_at
		)`,
		merchant); err != nil {
		tx.Rollback()
		helper.ErrorHandler(err, w)
		return
	}
	tx.Commit()
	helper.ReturnResponseAsJSON(w, nil, "Success Create User", 200)
}
