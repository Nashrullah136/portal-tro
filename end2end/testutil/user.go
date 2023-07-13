package testutil

import (
	"github.com/gavv/httpexpect/v2"
	"nashrul-be/crm/entities"
	"net/http"
	"time"
)

func CreateUser(e *httpexpect.Expect, req any) {
	auth := LoginAsAdmin(e)
	e.POST("/users").WithHeaders(auth).WithJSON(req).Expect().Status(http.StatusCreated)
	Logout(e, auth)
}

func Activate(username string) error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	user := entities.User{Username: username}
	_ = db.First(&user)
	*user.UpdatedAt = user.UpdatedAt.Add(time.Hour)
	return db.Table("users").Where(map[string]any{"username": user.Username}).Update("updated_at", user.UpdatedAt).Error
}
