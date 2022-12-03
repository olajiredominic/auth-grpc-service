package routes

import (
	"log"

	"gorm.io/gorm"

	models "github.com/lerryjay/auth-grpc-service/pkg/pb/model"
	"gorm.io/driver/postgres"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}

}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.UserORM{}, &models.AddressORM{}, &models.UserPermissionsORM{})

	return Handler{db}
}
