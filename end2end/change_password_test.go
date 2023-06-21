package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_change_password(t *testing.T) {
	testData, err := testutil.ReadYamlFile("change_password.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			if req, exist := data.Control["create"]; exist {
				testutil.CreateUser(e, req)
			}
			var auth map[string]string
			if req, exist := data.Control["login"]; exist {
				auth = testutil.Login(e, req)
			}
			responseBody := e.PATCH("/me").WithHeaders(auth).
				WithJSON(data.Data).Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				wantDataBefore := map[string]any{"password": "-"}
				wantDataAfter := map[string]any{"password": "-"}
				user := data.Expect["login"].(map[string]any)
				testutil.AssertAudit(t, user["username"].(string), "UPDATE", "USER",
					user["username"].(string), wantDataBefore, wantDataAfter)
				testutil.Login(e, data.Expect["login"])
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
