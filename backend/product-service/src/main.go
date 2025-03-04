package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"product-service/config"
	"product-service/controllers"
	"product-service/middlewares"
	"product-service/routes"

	_ "product-service/docs" // Import des docs Swagger

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/golang-jwt/jwt/v4"
)

// @title Product Service API
// @version 1.0
// @description API de gestion des produits pour la boutique en ligne
// @host localhost:4000
// @BasePath /
func main() {
	godotenv.Load()
	config.ConnectDB()

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Ajouter Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("🚀 Server running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
