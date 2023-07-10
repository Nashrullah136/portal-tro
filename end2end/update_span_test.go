package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"nashrul-be/crm/modules/span"
	"testing"
)

func Test_update_span(t *testing.T) {
	testData, err := testutil.ReadYamlFile("update_span.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			err = testutil.CreateSpan()
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
			responseBody := e.POST("/span/" + data.Data["documentNumber"].(string)).WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				oldSpan := testutil.GetSpan()
				newSpan := span.PatchBankRiau(oldSpan)
				dataBefore, _ := oldSpan.LogPresentation()
				dataAfter, _ := newSpan.LogPresentation()
				wantDataBefore := make(map[string]any)
				wantDataAfter := make(map[string]any)
				for key := range dataAfter {
					if dataBefore[key] != dataAfter[key] {
						wantDataBefore[key] = dataBefore[key]
						wantDataAfter[key] = dataAfter[key]
					}
				}
				testutil.AssertAudit(t, "user", "UPDATE", "SPAN",
					data.Data["documentNumber"].(string), wantDataBefore, wantDataAfter)
			}
			t.Cleanup(func() {
				err := testutil.DeleteSpan()
				if err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}
