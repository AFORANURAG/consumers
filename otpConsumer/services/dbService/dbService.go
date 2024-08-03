package dbservice

import (
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MYSQLDBService struct {
	db  *gorm.DB
}

func (dbs *MYSQLDBService) GetDb() (*gorm.DB, error) {
	return dbs.db, nil
}


var DbClient *MYSQLDBService
var once sync.Once

func NewDBServiceClientProvider(uri string) *MYSQLDBService {
	once.Do(func() {
		url:=(string)(uri)

		db, err := gorm.Open(mysql.Open(url),&gorm.Config{})
		if err != nil {
			log.Printf("Error while creating DBServiceClient:%v", err)
		}
		log.Println("Successfully connected to PlanetScale!")
		DbClient = &MYSQLDBService{db: db}
	})
	return DbClient
}
