package Database

import (
	"Exoplanet/Models"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbService interface {
	ExoplanetDb
	FuelDb
}

type Db struct {
	Data map[string]Models.Exoplanet
	Lock sync.RWMutex
	DB   *gorm.DB
}

func NewDBService() DbService {
	var db *gorm.DB
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Models.DbUser,
		Models.DbPassword,
		Models.DbHost,
		Models.DbPort,
		Models.DbName,
	)

	// Retry logic
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to connect to database after multiple attempts: %v", err)
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.Exoplanet{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	log.Println("database migration done")
	data := make(map[string]Models.Exoplanet, 0)
	return &Db{
		Data: data,
		DB:   db,
	}
}
