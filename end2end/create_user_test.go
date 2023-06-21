package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

// TODO: Add assert to database for audit log
func Test_create_user(t *testing.T) {
	testData, err := testutil.ReadYamlFile("create_user.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			var auth map[string]string
			loginAs, ok := data.Control["loginAs"].(string)
			if ok {
				switch loginAs {
				case "user":
					auth = testutil.LoginAsUser(e)
				case "admin":
					auth = testutil.LoginAsAdmin(e)
				}
			}
			responseBody := e.POST("/users").WithHeaders(auth).WithJSON(data.Data).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").Number().IsEqual(data.Expect["code"])
			responseBody.Value("message").NotNull().IsString().NotEqual("")
			if data.Control["case"].(string) == "success" {
				responseData := responseBody.Value("data").Object()
				responseData.Value("name").IsString().IsEqual(data.Expect["name"])
				responseData.Value("username").IsString().IsEqual(data.Expect["username"])
				responseData.Value("role").IsString().IsEqual(data.Expect["role"])
				responseData.Value("created_by").IsString().IsEqual("admin")
				responseData.Value("updated_by").IsString().IsEqual("admin")
				wantDataAfter := map[string]any{
					"password": "-",
					"username": data.Expect["username"],
					"name":     data.Expect["name"],
					"role_id":  2,
				}
				testutil.AssertAudit(t, data.Control["loginAs"].(string), "CREATE", "USER",
					data.Expect["username"].(string), nil, wantDataAfter)
				credential := map[string]string{
					"username": data.Data["username"].(string),
					"password": data.Data["password"].(string),
				}
				testutil.Login(e, credential)
			}
			t.Cleanup(func() {
				if data.Control["case"].(string) == "success" {
					db, _ := testutil.GetConn()
					db.Table("users").Where("username = ?", data.Data["username"].(string)).Delete(&map[string]any{})
				}
			})
		})
	}
}
