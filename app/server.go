package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adhiardiansyah/goshop/app/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.initializeDB(dbConfig)
	server.initializeRoutes()
	seeders.DBSeed(server.DB)
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBHost, dbConfig.DBName)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed on connecting to the database server")
	}

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrated successfully.")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoShop")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPass = getEnv("DB_PASS", "")
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBName = getEnv("DB_NAME", "goshop")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
