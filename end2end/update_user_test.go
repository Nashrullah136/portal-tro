package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"nashrul-be/crm/entities"
	"testing"
)

func Test_update_user(t *testing.T) {
	testData, err := testutil.ReadYamlFile("update_user.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			db, _ := testutil.GetConn()
			if err != nil {
				t.Fatal(err)
			}
			if req, exist := data.Control["create"]; exist {
				testutil.CreateUser(e, req)
			}
			var auth map[string]string
			if loginAs, ok := data.Control["loginAs"].(string); ok {
				auth = testutil.LoginAs(t, e, loginAs)
			}
			var createdUser entities.User
			if err := db.Where("username = ?", data.Control["create"].(map[string]any)["username"]).Find(&createdUser).Error; err != nil {
				t.Fatal(err)
			}
			responseBody := e.PATCH("/users/" + data.Data["username"].(string)).WithHeaders(auth).
				WithJSON(data.Data["update"]).Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				responseData := responseBody.Value("data").Object()
				responseData.Value("name").IsString().IsEqual(data.Expect["name"])
				responseData.Value("username").IsString().IsEqual(data.Expect["username"])
				responseData.Value("role").IsString().IsEqual(data.Expect["role"])
				responseData.Value("created_by").IsString().IsEqual("admin")
				responseData.Value("updated_by").IsString().IsEqual("admin")
				update := data.Data["update"].(map[string]any)
				wantDataBefore := make(map[string]any)
				wantDataAfter := make(map[string]any)
				if name, exist := update["name"]; exist {
					wantDataBefore["name"] = createdUser.Name
					wantDataAfter["name"] = name
				}
				if _, exist := update["password"]; exist {
					wantDataBefore["password"] = "-"
					wantDataAfter["password"] = "-"
				}
				testutil.AssertAudit(t, data.Control["loginAs"].(string), "UPDATE", "USER",
					data.Expect["username"].(string), wantDataBefore, wantDataAfter)
				testutil.Login(e, data.Expect["login"])
			}
			t.Cleanup(func() {
				req, exist := data.Control["create"]
				if exist {
					username := req.(map[string]any)["username"].(string)
					db.Table("users").Where("username = ?", username).Delete(&map[string]any{})
				}
			})
		})
	}
}
