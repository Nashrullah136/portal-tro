package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"nashrul-be/crm/entities"
	"net/http"
	"testing"
)

func Test_delete_user(t *testing.T) {
	db, _ := testutil.GetConn()
	testData, err := testutil.ReadYamlFile("delete_user.yaml")
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
			createdUser := entities.User{Username: data.Data["username"].(string)}
			if err := db.Find(&createdUser).Error; err != nil {
				t.Fatal(err)
			}
			e.DELETE("/users/" + data.Data["username"].(string)).WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int))
			if data.Control["case"].(string) == "success" {
				var audit entities.Audit
				if err := db.Order("date_time desc").First(&audit).Error; err != nil {
					t.Fatal(err)
				}
				wantDataBefore := map[string]any{
					"password": "-",
					"username": createdUser.Username,
					"name":     createdUser.Name,
					"role_id":  createdUser.RoleID,
				}
				testutil.AssertAudit(t, data.Control["loginAs"].(string), "DELETE", "USER",
					data.Data["username"].(string), wantDataBefore, nil)
				e.GET("/users/" + data.Data["username"].(string)).WithHeaders(auth).
					Expect().Status(http.StatusNotFound)
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
