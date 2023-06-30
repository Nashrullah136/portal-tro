package zabbix

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

type Server interface {
	Login() error
	Do(params any, result interface{}) error
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

func (z *server) Login() error {
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

// Do TODO: Handle when auth is not valid anymore
func (z *server) Do(params any, result interface{}) error {
	requestData := map[string]any{
		"jsonrpc": "2.0",
		"method":  "user.login",
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
	resultResponse := gjson.Get(string(respBody), "result")
	if err = json.Unmarshal([]byte(resultResponse.Raw), result); err != nil {
		return err
	}
	return nil
}
