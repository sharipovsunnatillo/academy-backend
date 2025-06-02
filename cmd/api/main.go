package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/middlewares"
	"github.com/spf13/viper"
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
