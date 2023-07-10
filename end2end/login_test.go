package end2end

import (
	"nashrul-be/crm/end2end/testutil"
	"testing"
)

func Test_login(t *testing.T) {
	testData, err := testutil.ReadYamlFile("login.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			e, err := testutil.InitTest(t)
			if err != nil {
				t.Fatal(err)
			}
			req := e.POST("/login").WithJSON(data.Data).Expect().Status(data.Expect["code"].(int))
			defer testutil.Logout(e, map[string]string{"Cookies": testutil.GetCookie(req)})
			req.Header("Set-Cookie").NotEmpty()
			responseBody := req.JSON().Object()
			responseBody.Value("code").Number().IsEqual(data.Expect["code"])
			responseBody.Value("message").NotNull().IsString().NotEqual("")
			responseData := responseBody.Value("data").Object()
			responseData.Value("role").IsString().IsEqual(data.Expect["role"])
			responseData.Value("username").IsString().IsEqual(data.Expect["username"])
			testutil.AssertAudit(t, data.Expect["username"].(string), "Login", "", "", nil, nil)
		})
	}
}
