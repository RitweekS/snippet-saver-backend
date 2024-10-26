package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDb() error {
	db_url := os.Getenv("DB_URL")
	dbConn,dbErr := gorm.Open(postgres.Open(db_url))
	if dbErr != nil {
		log.Printf("%#v\n DB_ERROR_CONNECTION\n", dbErr.Error())
		return  dbErr
	} else {
		log.Println("Connection Established")
	}

	DB = dbConn
	return dbErr	
}


func Close(){
	db,err := DB.DB()

	if err!=nil{
		log.Fatal("Failed to close database connection:", err)
	}

	db.Close()
}
