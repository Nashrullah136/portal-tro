package testutil

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
)

func CreateUser(e *httpexpect.Expect, req any) {
	auth := LoginAsAdmin(e)
	e.POST("/users").WithHeaders(auth).WithJSON(req).Expect().Status(http.StatusCreated)
	Logout(e, auth)
}
