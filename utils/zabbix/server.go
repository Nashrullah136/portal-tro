package zabbix

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"nashrul-be/crm/utils/logutils"
	"net/http"
)

type Server interface {
	Login() error
	Do(method string, params any, result interface{}) error
}

type server struct {
	url      string
	username string
	password string
	auth     string
}

func NewServer(url, username, password string) Server {
	return &server{
		url:      url,
		username: username,
		password: password,
	}
}

func (z *server) isLogin() bool {
	return z.auth != ""
}

func (z *server) Login() error {
	logutils.Get().Println("Trying login to zabbix server...")
	if z.auth != "" {
		return nil
	}
	request := map[string]any{
		"jsonrpc": "2.0",
		"method":  "user.login",
		"params": map[string]any{
			"user":     z.username,
			"password": z.password,
		},
		"id": 1,
	}
	reqOdy, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := http.Post(z.url, "application/json", bytes.NewBuffer(reqOdy))
	if err != nil {
		return err
	}
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	z.auth = gjson.Get(string(respBody), "result").String()
	return nil
}

func (z *server) Do(method string, params any, result interface{}) error {
	if !z.isLogin() {
		if err := z.Login(); err != nil {
			return err
		}
	}
	requestData := map[string]any{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"auth":    z.auth,
		"id":      1,
	}
	reqBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	response, err := http.Post(z.url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	errObj := gjson.Get(string(respBody), "error.data")
	if errObj.String() == "Session terminated, re-login, please." {
		if err = z.Login(); err != nil {
			logutils.Get().Println("Failed to re-login to zabbix server")
			return err
		}
		response, err = http.Post(z.url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}
	}
	resultResponse := gjson.Get(string(respBody), "result")
	if err = json.Unmarshal([]byte(resultResponse.Raw), result); err != nil {
		return err
	}
	return nil
}
