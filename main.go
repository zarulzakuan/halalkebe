package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	HalalStatus  string `json:"status"`
	Ingredients  string `json:"ingredients"`
	PhotoPath    string `json:"photopath"`
	Submitter    string `json:"submitter"`
}

var products []Product

func ourInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Product{})
}
func createEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(&product)
}
func getEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	barcodeNum := params["id"]

	mockupProd := &Product{
		ID:           barcodeNum,
		Name:         "Kara Mucho Spicy Potato Chips",
		Manufacturer: "KOIKEYA",
		HalalStatus:  "Y",
		Ingredients:  "potatoes, palm oil, sugar, dextrose, monosodium glutamate, salt, seaweed, soy sauce powder (soybeans, wheat), hydrolyzed soy protein, white pepper, yeast extract, sesame seed oil, onion powder, natural and artificial flavor, tricalcium phosphate, spices extract, silicon dioxide, DL-Methionine, disodium",
		PhotoPath:    "https://static.openfoodfacts.org/images/products/489/705/364/0026/front_en.5.full.jpg",
		Submitter:    "zarulzakuan@gmail.com",
	}
	log.Println(mockupProd)
	json.NewEncoder(w).Encode(mockupProd)
}
func updateEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	barcodeNum := params["id"]
	json.NewEncoder(w).Encode(&Product{ID: barcodeNum})
}
func deleteEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	barcodeNum := params["id"]
	json.NewEncoder(w).Encode(&Product{ID: barcodeNum})
}
func main() {
	router := mux.NewRouter()

	//router.HandleFunc("/", ourInfo).Methods("GET")
	//router.HandleFunc("/entry", createEntry).Methods("POST")
	router.HandleFunc("/entry/{id}", getEntry).Methods("GET")
	//router.HandleFunc("/entry/{id}", updateEntry).Methods("PUT")
	//router.HandleFunc("/entry/{id}", deleteEntry).Methods("DELETE")
	log.Println("Starting...")
	http.ListenAndServe(":8000", router)
}
