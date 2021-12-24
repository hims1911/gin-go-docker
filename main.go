package main

import (
	"log"
	"net/http"

	buisness "go-gin-docker/buisness"
	models "go-gin-docker/models"

	gin "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	// fetching the account details
	buisness.FetchCondition()
	// GIN Router Code
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
	db.AutoMigrate(&models.Product{})

	// Get Product list route
	r.GET("/products", func(c *gin.Context) {
		// db find will fetch the table value and will assign it to the product
		var products []models.Product
		if err := db.Find(&products).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			log.Println(err)
			return
		} else {
			log.Println(products)
			c.JSON(http.StatusOK, products)
		}

		// read from aws dynamo db
	})

	// Get Product:ID route
	r.GET("/products/:id", func(context *gin.Context) {
		// param will help to take the parameter
		id := context.Param("id")
		log.Println(id)

		// db.Find can be passed with id as condition it act as exist
		var products models.Product
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
		var createProduct models.Product
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
		var product models.Product
		if err := db.Find(&product, id).Error; err != nil {
			log.Println("ID not Exist", err)
			context.JSON(http.StatusBadRequest, gin.H{"error": "Records Not Found"})
			return
		}

		// assign the body values to the Update Struct Model
		var createProduct models.Product
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
		var product models.Product
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
