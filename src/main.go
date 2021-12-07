package main

import (
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	ID    uint   `form:"id" json:"id"`
	Code  string `form:"code" json:"code"`
	Price uint   `form:"price" json:"price"`
}

type CreateProduct struct {
	Code  string `form:"code" json:"code"`
	Price uint   `form:"price" json:"price"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Got the data",
		})
	})

	log.Println("Starting the Server")

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		log.Println("Failed to Connect the Database")
		panic("Failed to Connect Databse")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	// Get Product list route
	r.GET("/products", func(c *gin.Context) {
		// db find will fetch the table value and will assign it to the product
		var products []Product
		if err := db.Find(&products).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		} else {
			log.Println(products)
			c.JSON(http.StatusOK, products)
		}
	})

	// Get Product:ID route
	r.GET("/products/:id", func(context *gin.Context) {
		// param will help to take the parameter
		id := context.Param("id")
		log.Println(id)

		// db.Find can be passed with id as condition it act as exist
		var products Product
		if err := db.Find(&products, id).Error; err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			log.Println("got the ID")
			context.JSON(http.StatusOK, products)
		}
	})

	// POST request to create product route
	r.POST("/createproduct", func(context *gin.Context) {
		// attaching the product with the Query Parameter
		var createProduct Product
		if err := context.ShouldBind(&createProduct); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// creating the variable that binds the product
		// returning the response
		db.Create(&createProduct)

		context.JSON(http.StatusOK, gin.H{"data": createProduct})
	})

	// Patch the Request
	r.PATCH("products/:id", func(context *gin.Context) {
		// Fetching the ID
		id := context.Param("id")
		log.Println(id)

		// Checking If ID exists or Not
		var product Product
		if err := db.Find(&product, id).Error; err != nil {
			log.Println("ID not Exist", err)
			context.JSON(http.StatusBadRequest, gin.H{"error": "Records Not Found"})
			return
		}

		// assign the body values to the Update Struct Model
		var createProduct CreateProduct
		if err := context.ShouldBindJSON(&createProduct); err != nil {
			log.Println(createProduct)
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update the record with the passed record struct
		db.Find(&product).Updates(createProduct)

		context.JSON(http.StatusOK, gin.H{"data": createProduct})
	})

	// Delete Route: Delete the Product
	r.DELETE("/products/:id", func(context *gin.Context) {
		// finding the ID from the Parameter and checking if Table Exist or Not
		id := context.Param("id")
		var product Product
		if err := db.Find(&product, id).Error; err != nil {
			log.Println("Id not found", err)
			context.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found"})
			return
		}

		// Deleting the Modal
		if err := db.Delete(&product).Error; err != nil {
			log.Println("Error while Deleting", err)
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": "Item Deleted Successfully"})
	})

	r.Run(":8080")
}
