package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go-rest-api/internal/database"
	"go-rest-api/internal/user"
)

func main() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	userRepo := user.NewRepository(db)
	userHandler := user.NewHandler(userRepo)

	user.RegisterRoutes(r, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}