package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_get_profile(t *testing.T) {
	testData, err := testutil.ReadYamlFile("get_profile.yaml")
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
			responseBody := e.GET("/me").WithHeaders(auth).WithJSON(data.Data).
				Expect().Status(data.Expect["code"].(int)).JSON().Object()
			responseBody.Value("code").Number().IsEqual(data.Expect["code"])
			responseBody.Value("message").NotNull().IsString().NotEqual("")
			if data.Control["case"].(string) == "success" {
				responseData := responseBody.Value("data").Object()
				responseData.Value("name").IsString().IsEqual(data.Expect["name"])
				responseData.Value("username").IsString().IsEqual(data.Expect["username"])
				responseData.Value("role").IsString().IsEqual(data.Expect["role"])
				responseData.Value("created_by").IsString().IsEqual(data.Expect["createdBy"])
				responseData.Value("updated_by").IsString().IsEqual(data.Expect["updatedBy"])
			}
		})
	}
}
