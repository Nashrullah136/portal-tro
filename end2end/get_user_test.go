package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_get_user(t *testing.T) {
	testData, err := testutil.ReadYamlFile("get_user.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			req, exist := data.Control["create"]
			if exist {
				testutil.CreateUser(e, req)
			}
			var auth map[string]string
			if loginAs, ok := data.Control["loginAs"].(string); ok {
				auth = testutil.LoginAs(t, e, loginAs)
			}
			defer testutil.Logout(e, auth)
			responseBody := e.GET("/users/" + data.Data["username"].(string)).WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				responseData := responseBody.Value("data").Object()
				responseData.Value("name").IsString().IsEqual(data.Expect["name"])
				responseData.Value("username").IsString().IsEqual(data.Expect["username"])
				responseData.Value("role").IsString().IsEqual(data.Expect["role"])
				responseData.Value("created_by").IsString().IsEqual("admin")
				responseData.Value("updated_by").IsString().IsEqual("admin")
			}
			t.Cleanup(func() {
				req, exist := data.Control["create"]
				if exist {
					username := req.(map[string]any)["username"].(string)
					db, _ := testutil.GetConn()
					db.Table("users").Where("username = ?", username).Delete(&map[string]any{})
				}
			})
		})
	}
}
