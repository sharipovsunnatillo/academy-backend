package main

import (
	"database/sql"
	"github.com/sharipov/sunnatillo/academy-backend/internal/database"
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/enums"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func main() {
	db, err := database.NewDB(database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "123456",
		DBName:   "postgres",
	})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute)

	err = db.AutoMigrate(&models.User{}, models.Subject{})
	if err != nil {
		log.Fatalln(err)
	}

	roles := []models.Role{{Name: string(enums.SUPER_ADMIN)}, {Name: string(enums.ADMIN)}, {Name: string(enums.TEACHER)}, {Name: string(enums.STUDENT)}, {Name: string(enums.GUEST)}, {Name: string(enums.PARENT)}}
	tx := db.Clauses(clause.OnConflict{DoNothing: true}).Create(roles)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	permissions := []models.Permission{{Name: "create_user"}, {Name: "update_user"}, {Name: "delete_user"}, {Name: "create_subject"}, {Name: "update_subject"}, {Name: "delete_subject"}}
	tx = db.Clauses(clause.OnConflict{DoNothing: true}).Create(permissions)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	user := models.User{
		Active:      false,
		FirstName:   "Sunnatillo",
		LastName:    "Sharipov",
		MiddleName:  "Jamshid o'g'li",
		Gender:      sql.Null[enums.Gender]{V: enums.Male, Valid: true},
		Birthday:    sql.NullTime{Time: time.Now().Add(-time.Hour * 24 * 365 * 25), Valid: true},
		Phone:       "09999999999",
		Password:    "123456789",
		Roles:       []*models.Role{&roles[0]},
		Permissions: []*models.Permission{&permissions[0]},
	}
	tx = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
}
