package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_logout(t *testing.T) {
	testData, err := testutil.ReadYamlFile("logout.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			auth := make(map[string]string)
			loginAs, ok := data.Control["loginAs"].(string)
			switch {
			case len(data.Data) > 0:
				for key, val := range data.Data {
					auth[key] = val.(string)
				}
			case ok:
				auth = testutil.LoginAs(t, e, loginAs)
			}
			req := e.GET("/logout").WithHeaders(auth).Expect().Status(data.Expect["code"].(int))
			responseBody := req.JSON().Object()
			responseBody.Value("code").Number().IsEqual(data.Expect["code"])
			if val, exist := data.Control["withCookie"]; exist && val == "true" {
				req.Header("Set-Cookie").NotEmpty()
			}
			if data.Control["case"].(string) == "success" {
				testutil.AssertAudit(t, data.Expect["username"].(string), "Logout", "", "", nil, nil)
			}
		})
	}
}
