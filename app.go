package main

import (
	"log"

	"github.com/DazEdword/go-gin-dockerised/db"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Autoloading .env (only when running undockerised!)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start Db
	var connectionString = db.BuildConnectionString()
	db.InitDb(connectionString)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
