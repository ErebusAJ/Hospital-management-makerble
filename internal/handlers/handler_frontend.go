package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loginHTML(c *gin.Context) {
	godotenv.Load()

	api := os.Getenv("API")
	c.HTML(200, "login.html", gin.H{
		"addr" : api, 
	})
}