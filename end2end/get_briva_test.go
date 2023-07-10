package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_get_briva(t *testing.T) {
	testData, err := testutil.ReadYamlFile("get_briva.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			err = testutil.CreateBriva()
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
			defer testutil.Logout(e, auth)
			responseBody := e.GET("/briva/" + data.Data["brivano"].(string)).WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				responseData := responseBody.Value("data").Object()
				responseData.Value("Brivano").IsString().IsEqual(data.Expect["brivano"])
			}
			t.Cleanup(func() {
				err := testutil.DeleteBriva()
				if err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
