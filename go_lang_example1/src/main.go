package main

import (
	"fmt"
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	ID    uint   `json:"id"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Got the data",
		})
	})

	log.Println("Starting the Server")

	r.GET("/products", func(c *gin.Context) {
		db, err := gorm.Open("sqlite3", "test.db")

		if err != nil {
			log.Println("Failed to Connect the Database")
			panic("Failed to Connect Databse")
		}

		defer db.Close()

		db.AutoMigrate(&Product{})

		var products []Product

		if err := db.Find(&products).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
		} else {
			fmt.Println(products)
			c.JSON(http.StatusOK, products)
			log.Println("Products Returned")
		}
	})

	r.Run(":8080")
}
