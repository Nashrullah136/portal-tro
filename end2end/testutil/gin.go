package testutil

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nashrul-be/crm/app"
)

func SetUpGin(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	app.Handle(db, engine)
	return engine
}
