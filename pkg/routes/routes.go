package routes

import (
	"log"

	"gorm.io/gorm"

	models "github.com/lerryjay/auth-grpc-service/pkg/pb/model"
	"gorm.io/driver/postgres"
)

type Handler struct {
	DB                     *gorm.DB
	ClientID               string
	SecretKey              string
	TokenURL               string
	QoreidBaseURL          string
	BiometricQoreidBaseURL string
	VNINURL                string
	NINURL                 string
	DLURL                  string
	PassportURL            string
}

// New creates a new Handler with the provided database connection
func New(db *gorm.DB, clientID, secretKey, tokenURL, qoreidBaseURL, vninURL, ninURL, dlURL, passportURL, BiometricQoreidBaseURL string) Handler {
	return Handler{
		DB:                     db,
		ClientID:               clientID,
		SecretKey:              secretKey,
		TokenURL:               tokenURL,
		QoreidBaseURL:          qoreidBaseURL,
		VNINURL:                vninURL,
		NINURL:                 ninURL,
		DLURL:                  dlURL,
		PassportURL:            passportURL,
		BiometricQoreidBaseURL: BiometricQoreidBaseURL,
	}
}

// Init initializes the database connection and returns a Handler
func Init(url, clientID, secretKey, tokenURL, qoreidBaseURL, vninURL, ninURL, dlURL, passportURL, BiometricQoreidBaseURL string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.UserORM{}, &models.UserVerificationORM{}, &models.AddressORM{}, &models.UserPermissionORM{})

	return Handler{
		DB:                     db,
		ClientID:               clientID,
		SecretKey:              secretKey,
		TokenURL:               tokenURL,
		QoreidBaseURL:          qoreidBaseURL,
		VNINURL:                vninURL,
		NINURL:                 ninURL,
		DLURL:                  dlURL,
		PassportURL:            passportURL,
		BiometricQoreidBaseURL: BiometricQoreidBaseURL,
	}
}
