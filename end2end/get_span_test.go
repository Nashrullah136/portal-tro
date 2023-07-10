package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_get_span(t *testing.T) {
	testData, err := testutil.ReadYamlFile("get_span.yaml")
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
			if loginAs, ok := data.Control["loginAs"].(string); ok {
				auth = testutil.LoginAs(t, e, loginAs)
			}
			defer testutil.Logout(e, auth)
			responseBody := e.GET("/span/" + data.Data["documentNumber"].(string)).WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
			if data.Control["case"].(string) == "success" {
				responseData := responseBody.Value("data").Object()
				responseData.Value("DocumentNumber").IsString().IsEqual(data.Expect["DocumentNumber"])
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
