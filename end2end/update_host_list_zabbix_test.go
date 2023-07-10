package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_update_host_list_zabbix(t *testing.T) {
	testData, err := testutil.ReadYamlFile("update_host_list_zabbix.yaml")
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
			if loginAs, ok := data.Control["loginAs"].(string); ok {
				auth = testutil.LoginAs(t, e, loginAs)
			}
			defer testutil.Logout(e, auth)
			responseBody := e.GET("/server-utilization/update-host").WithHeaders(auth).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").IsNumber().IsEqual(data.Expect["code"])
		})
	}
}
