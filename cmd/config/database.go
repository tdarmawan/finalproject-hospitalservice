package config

import (
	"fmt"
	"log"
	"time"

	// "database/sql"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"testing/cmd/services"
)

type Database struct {
	Name     string `env:"DB_SCHEMA"`
	Adapter  string `env:"DB_DRIVER"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	UserDB   string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type ServeConfig struct {
	ServiceName string `env:"SERVICE_NAME"`
	ServicePort string `env:"SERVICE_PORT"`
	ServiceHost string `env:"SERVICE_HOST"`
	db          Database
}

var db *gorm.DB

var Config ServeConfig

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Cannot find .env file. OS Environments will be used")
	}

	err = env.Parse(&Config)
	if err != nil {
		log.Fatal("Web Server not working")
	}

	err = env.Parse(&Config.db)

	return err

}

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}

	ConnectDB()
}

func ConnectDB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbConfig := Config.db
	if dbConfig.Adapter == "postgres" {
		config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.UserDB, dbConfig.Password, dbConfig.Name, dbConfig.Port)
		dsn := config
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("Connected to Database Local postgressql")
	}

	if err != nil {
		log.Println("[Driver.ConnectDB] error when connect to database")
		log.Fatal("[Driver.ConnectDB] error when connect to database")
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	pingDb := func(db *gorm.DB) {
		sqldb.Ping()
	}

	go doEvery(6*time.Minute, pingDb, db)

	// Database Pooling
	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(200)
	sqldb.SetConnMaxLifetime(45 * time.Second)

	db.Debug().AutoMigrate(&services.User{}, &services.Dokter{}, &services.Hospital{}, &services.Schedule{}, &services.Appointment{}, &services.Patient{}, &services.HospitalPatient{})

	return db
}

func doEvery(d time.Duration, f func(*gorm.DB), y *gorm.DB) {
	for _ = range time.Tick(d) {
		f(y)
	}
}
