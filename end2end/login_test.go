package end2end

import (
	"github.com/stretchr/testify/assert"
	"nashrul-be/crm/end2end/testutil"
	"nashrul-be/crm/entities"
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
			req.Header("Set-Cookie").NotEmpty()
			responseBody := req.JSON().Object()
			responseBody.Value("code").Number().IsEqual(data.Expect["code"])
			responseBody.Value("message").NotNull().IsString().NotEqual("")
			responseData := responseBody.Value("data").Object()
			responseData.Value("role").IsString().IsEqual(data.Expect["role"])
			responseData.Value("username").IsString().IsEqual(data.Expect["username"])
			db, err := testutil.GetConn()
			if err != nil {
				t.Fatal(err)
			}
			var audit entities.Audit
			if err := db.Order("date_time desc").First(&audit).Error; err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, data.Expect["username"], audit.Username)
			assert.Equal(t, "Login", audit.Action)
		})
	}
}
