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
	return map[string]string{
		"Cookie": GetCookie(req),
	}
}

func GetCookie(response *httpexpect.Response) string {
	setCookie := response.Header("Set-Cookie").Raw()
	return strings.SplitN(setCookie, ";", 2)[0]
}

func Logout(e *httpexpect.Expect, credential map[string]string) {
	e.GET("/logout").WithHeaders(credential)
}
