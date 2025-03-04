package controllers

import (
	"encoding/json"
	"net/http"
	"product-service/config"
	"product-service/models"
	"strconv"

	"github.com/gorilla/mux"
)

// 🔹 Ajouter un produit
// @Summary Ajouter un produit
// @Description Ajoute un nouveau produit
// @Tags Products
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Détails du produit"
// @Success 201 {object} models.Product
// @Router /api/products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	query := `INSERT INTO products (name, price, stock, category) VALUES ($1, $2, $3, $4) RETURNING id`
	err := config.DB.QueryRow(query, product.Name, product.Price, product.Stock, product.Category).Scan(&product.ID)
	if err != nil {
		http.Error(w, "Erreur d'insertion", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// 🔹 Récupérer tous les produits
// @Summary Récupérer tous les produits
// @Description Retourne la liste de tous les produits
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Product
// @Router /api/products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, price, stock, category FROM products")
	if err != nil {
		http.Error(w, "Erreur de récupération", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Category)
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}

// 🔹 Supprimer un produit
// @Summary Supprimer un produit
// @Description Supprime un produit par son ID
// @Tags Products
// @Param id path int true "ID du produit"
// @Success 204
// @Router /api/products/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Erreur suppression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
