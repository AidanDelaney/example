package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	hello := Message{
		Message: "hello world",
	}
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true;
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, hello)
	})

	r.Run()
}
