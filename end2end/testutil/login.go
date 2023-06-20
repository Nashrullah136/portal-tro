package testutil

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"strings"
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
	req := e.POST("/login").WithJSON(credential).Expect().Status(http.StatusOK)
	setCookie := req.Header("Set-Cookie").Raw()
	cookie := strings.SplitN(setCookie, ";", 2)[0]
	return map[string]string{
		"Cookie": cookie,
	}
}
