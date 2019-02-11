package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	database "./database"
	"github.com/gorilla/mux"
)

var db database.IDataBase = new(database.DataBase)

func getResponseMap(result interface{}, err error) map[string]interface{} {

	response := make(map[string]interface{})

	if err == nil {
		response["response"] = result
		response["err"] = ""
	} else {
		response["response"] = nil
		response["err"] = err
	}

	return response
}

func getSortedProductsByCost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := db.GetAll()

	json.NewEncoder(w).Encode(getResponseMap(products, err))
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		json.NewEncoder(w).Encode(getResponseMap(nil, err))
		return
	}

	product, err := db.Get(uint64(id))
	json.NewEncoder(w).Encode(getResponseMap(product, err))
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product database.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	id, err := db.AddProduct(product)
	json.NewEncoder(w).Encode(getResponseMap(map[string]uint64{"id": id}, err))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		json.NewEncoder(w).Encode(getResponseMap(nil, err))
		return
	}

	var product database.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	err = db.Change(product)
	json.NewEncoder(w).Encode(getResponseMap(product, err))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		json.NewEncoder(w).Encode(getResponseMap(nil, err))
		return
	}

	err = db.DeleteProduct(uint64(id))
	json.NewEncoder(w).Encode(getResponseMap(true, err))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/product/sort", getSortedProductsByCost).Methods("GET")
	r.HandleFunc("/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/pruduct", addProduct).Methods("POST")
	r.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
