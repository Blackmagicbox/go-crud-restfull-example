package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Products []Product

func allProducts(w http.ResponseWriter, r *http.Request) {
	products := Products{
		Product{
			Id:    0,
			Name:  "iphone-x",
			Price: 1990.00,
		},
	}

	fmt.Println("Products Endpoint hit")
	json.NewEncoder(w).Encode(products)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/products", allProducts)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
