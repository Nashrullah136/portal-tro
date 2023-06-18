package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nashrul-be/crm/app"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/translate"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	if err := translate.RegisterTranslator(); err != nil {
		panic(err.Error())
	}

	engine := gin.Default()

	dbConn, err := db.Connect(db.DsnWithEnv())
	if err != nil {
		panic(err.Error())
	}

	app.Handle(dbConn, engine)

	if err = engine.Run(); err != nil {
		panic(err.Error())
	}
}
