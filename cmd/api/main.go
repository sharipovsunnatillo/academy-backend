package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/sharipov/sunnatillo/academy-backend/internal/models"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/middlewares"
	"github.com/sharipov/sunnatillo/academy-backend/seed"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	env := flag.String("env", "dev", "environment")
	flag.Parse()

	viper.SetConfigName(*env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("❌ Error reading config file, %s", err)
	}

	httpPort := viper.GetInt("http.port")

	middleware := middlewares.Ensure(middlewares.Logging)
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: middleware(mux),
	}

	fmt.Println("==============================================================================================")
	fmt.Println("\t\tApp started successfully")
	fmt.Println("\t\tEnvironment: ", *env)
	fmt.Println("\t\tListening on port ", server.Addr)
	fmt.Println("==============================================================================================")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC",
		viper.GetString("datasource.host"),
		viper.GetInt("datasource.port"),
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.database"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatalf("❌ Failed to get tables: %v", err)
	}
	for _, table := range tables {
		err := db.Migrator().DropTable(table)
		if err != nil {
			log.Fatalf("❌ Failed to drop table %s: %v", table, err)
		}
	}

	err = db.AutoMigrate(
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

	seed.Populate(db)

	// Run server in goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	//Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop // wait for signal
	log.Println("⏳ Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Forced shutdown: %v", err)
	}

	log.Println("✅ Server exited cleanly")
}
