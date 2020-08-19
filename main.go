package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	product := &Product{
		ProductID:   123,
		ProductName: "pencil",
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(productJSON))
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
