package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_create_audit(t *testing.T) {
	testData, err := testutil.ReadYamlFile("create_audit.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			db, _ := testutil.GetConn()
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			if req, exist := data.Control["create"]; exist {
				testutil.CreateUser(e, req)
			}
			var auth map[string]string
			if req, exist := data.Data["login"]; exist {
				auth = testutil.Login(e, req)
			}
			responseBody := e.POST("/audits").WithHeaders(auth).
				WithJSON(data.Data["req"]).Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				user := data.Data["login"].(map[string]any)
				req := data.Data["req"].(map[string]any)
				testutil.AssertAudit(t, user["username"].(string), req["action"].(string), "", "", nil, nil)
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
