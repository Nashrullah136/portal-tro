package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"net/http"
	"testing"
)

func Test_delete_user(t *testing.T) {
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
			loginAs, ok := data.Control["loginAs"].(string)
			if ok {
				switch loginAs {
				case "user":
					auth = testutil.LoginAsUser(e)
				case "admin":
					auth = testutil.LoginAsAdmin(e)
				}
			}
			e.DELETE("/users/" + data.Data["username"].(string)).WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int))
			if data.Control["case"].(string) == "success" {
				e.GET("/users/" + data.Data["username"].(string)).WithHeaders(auth).
					Expect().Status(http.StatusNotFound)
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
