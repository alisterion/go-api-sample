package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sample/pkg/app"
	"sample/pkg/models"
	"sample/pkg/routers"
)


func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	InitConfig()
	app.InitDb()

	defer app.Db.Close()
	app.Db.AutoMigrate(&models.User{})

	gin.SetMode(gin.DebugMode)

	router := routers.SetupRouter()
	_ = router.Run(os.Getenv("SERVER_PORT"))
}
