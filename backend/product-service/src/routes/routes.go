package routes

import (
	"net/http"
	"product-service/controllers"
	"product-service/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	protected := router.PathPrefix("/api/admin").Subrouter()
	protected.Use(middlewares.JWTAuthMiddleware)
	protected.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	// Ajouter Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("🚀 Server running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
	return router
}
