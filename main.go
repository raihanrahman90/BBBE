// main.go
package main

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		err = godotenv.Load()
		fmt.Print("error loading .env.development file")
	}

	// Initialize the database
	config.ConnectDatabase()

	// Auto migrate the User and Teacher tables
	config.DB.AutoMigrate(&models.User{}, &models.LandingPage{}, &models.History{}, &models.Article{}, &models.Testimoni{}, &models.Order{}, &models.OrderItem{}, &models.Item{}, &models.Address{}, &models.Cart{})

	// Initialize the router
	r := routes.SetupRouter()

	// Start the server
	port := os.Getenv("PORT")
	fmt.Sprint("server started at :", port)
	fmt.Println(os.Getenv("PATH_STATIC"))
	fmt.Println(err)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
