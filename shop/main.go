package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	_AlgoHttp "shop/algorithm/delivery/http"
	_AlgoRepo "shop/algorithm/repository/rest"
	_AlgoUsecase "shop/algorithm/usecase"
	_PurchHttp "shop/purchase/delivery/http"
	_PurchRepo "shop/purchase/repository/rest"
	_PurchUsecase "shop/purchase/usecase"
	_userHttp "shop/user/delivery/http"
	_userRepo "shop/user/repository/postgres"
	_userUsecase "shop/user/usecase"

	"syscall"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// Construct the connection string
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	fmt.Println(connection)
	// Open a connection to the database
	dbConn, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// Close the database connection
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	g := gin.Default()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	AlgoRepoRest := _AlgoRepo.NewAlgorithmRepository("http://127.0.0.1:3000")
	AlgoUcase := _AlgoUsecase.NewUserUsecase(AlgoRepoRest, timeoutContext)
	_AlgoHttp.NewAlgoHandler(g, AlgoUcase)

	userRepo := _userRepo.NewUserRepository(dbConn)
	userUcase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	_userHttp.NewUserHandler(g, userUcase)

	PurchRepoRest := _PurchRepo.NewPurchaseRepository("http://127.0.0.1:3000")
	PurchUcase := _PurchUsecase.NewPurchaseUsecase(PurchRepoRest, userRepo, timeoutContext)
	_PurchHttp.NewPurchaseHandler(g, PurchUcase)

	server := &http.Server{
		Addr:    viper.GetString("server.address"),
		Handler: g,
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for a termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
