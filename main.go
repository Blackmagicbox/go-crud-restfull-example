package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Products []Product

var productsList = listProducts()

func listProducts() Products {
	products := Products{
		Product{
			Id:    0,
			Name:  "iphone-x",
			Price: 1990.00,
		},
		Product{
			Id:    2,
			Name:  "ipad Pro",
			Price: 950.00,
		},
	}
	return products
}

func allProducts(c *gin.Context) {
	fmt.Println("Products Endpoint hit")
	c.JSON(http.StatusOK, productsList)
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "Greetings")
}

func handleRequest() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/products", allProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", addProduct)
	_ = router.Run(":8081")
}

func addProduct(c *gin.Context) {

	var (
		product Product
	)
	_ = c.BindJSON(&product)

	productsList = append(productsList, product)
	c.JSON(http.StatusOK, gin.H{
		"price": product.Price,
		"id": product.Id,
		"name": product.Name,
	})


}

func getProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	products := productsList
	for _, product := range products {
		if product.Id == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusBadRequest, Products{})
}

func main() {
	handleRequest()
}
