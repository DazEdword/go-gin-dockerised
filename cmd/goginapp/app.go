package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/DazEdword/go-gin-dockerised/db"

	"github.com/gin-gonic/gin"
)

func main() {
	pwd, _ := os.Getwd()

	r := CreateApp(pwd)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// CreateApp is the application factory
func CreateApp(rootPath string) *gin.Engine {
	// Start Db
	var connectionString = db.BuildConnectionString()
	db.InitDb(connectionString)

	r := gin.Default()

	templatesPath := filepath.Join(rootPath, "templates")

	// Loading templates
	r.LoadHTMLGlob(templatesPath + "/*")

	// Main index
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello world!",
		})
	})

	// Defining routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": GetUsers(),
		})
	})

	return r
}

// GetUsers query the "users" table and retrieve a slice of JSON objects
func GetUsers() []interface{} {
	// Query json directly
	query := "SELECT row_to_json(r) FROM (SELECT id, username, email from users) r;"
	rows, err := db.Db.Query(context.Background(), query)

	if err != nil {
		log.Panicf("QueryRow failed: %v\n", err)
	}

	var jsonData []interface{}

	for rows.Next() {
		val, err := rows.Values()
		if err != nil {
			log.Panic("Error: Could not retrieve users", err)
		}

		jsonData = append(jsonData, val[0])
	}

	return jsonData
}
