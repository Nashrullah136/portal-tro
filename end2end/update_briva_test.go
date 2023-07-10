package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_update_briva(t *testing.T) {
	testData, err := testutil.ReadYamlFile("update_briva.yaml")
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
			responseBody := e.POST("/briva/" + data.Data["brivano"].(string)).WithHeaders(auth).
				WithJSON(data.Data["update"]).Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				oldBriva := testutil.GetBriva()
				newBriva := oldBriva
				newBriva.IsActive = data.Data["update"].(map[string]any)["active"].(string)
				dataBefore, _ := oldBriva.LogPresentation()
				dataAfter, _ := newBriva.LogPresentation()
				wantDataBefore := make(map[string]any)
				wantDataAfter := make(map[string]any)
				for key := range dataAfter {
					if dataBefore[key] != dataAfter[key] {
						wantDataBefore[key] = dataBefore[key]
						wantDataAfter[key] = dataAfter[key]
					}
				}
				testutil.AssertAudit(t, "user", "UPDATE", "BRIVA",
					data.Data["brivano"].(string), wantDataBefore, wantDataAfter)
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
