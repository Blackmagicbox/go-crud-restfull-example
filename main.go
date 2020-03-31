package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Products []Product

func allProducts(c *gin.Context) {
	products := Products{
		Product{
			Id:    0,
			Name:  "iphone-x",
			Price: 1990.00,
		},
	}

	fmt.Println("Products Endpoint hit")
	c.JSON(http.StatusOK, products)
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "Greetings")
}

func handleRequest() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/products", allProducts)
	router.Run(":8081")
}

func main() {
	handleRequest()
}
