package database

import (
	"fmt"
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
}

func NewDB(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tashkent", config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			LogLevel: logger.Info,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	//tables, err := db.Migrator().GetTables()
	//if err != nil {
	//	log.Fatalf("❌ Failed to get tables: %v", err)
	//}
	//for _, table := range tables {
	//	err := db.Migrator().DropTable(table)
	//	if err != nil {
	//		log.Fatalf("❌ Failed to drop table %s: %v", table, err)
	//	}
	//}

	err := db.AutoMigrate(
		&models.Region{},
		&models.District{},
		&models.Role{},
		&models.Permission{},
		&models.Subject{},
		&models.TextBook{},
		&models.TimeSlot{},
		&models.TrainingCenter{},
		&models.Branch{},
		&models.Room{},
		&models.User{},
		&models.TeacherInfo{},
		&models.Document{},
		&models.Group{},
		&models.Lesson{},
		&models.Task{},
		&models.Attendance{},
		&models.Grade{},
	)

	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	//seed.Populate(db) //todo
}
