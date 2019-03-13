package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var Db *gorm.DB

func InitDb() {
	DbName := os.Getenv("DB_NAME")
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dbConf := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", DbHost, DbPort, DbName, DbUser, DbPass)
	Db, _ = gorm.Open("postgres", dbConf)
}
