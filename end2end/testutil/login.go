package testutil

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
)

func LoginAsAdmin(e *httpexpect.Expect) map[string]string {
	credential := map[string]string{
		"username": "admin",
		"password": "admin",
	}
	return Login(e, credential)
}

func LoginAsUser(e *httpexpect.Expect) map[string]string {
	credential := map[string]string{
		"username": "user",
		"password": "user",
	}
	return Login(e, credential)
}

func Login(e *httpexpect.Expect, credential any) map[string]string {
	body := e.POST("/login").WithJSON(credential).Expect().Status(http.StatusOK).JSON().Object()
	token := body.Value("data").Object().Value("token").String().Raw()
	return map[string]string{
		"Authorization": token,
	}
}
