package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Custom Error function
func ErrorJSON(c *gin.Context, code int, client, server string, err error){
	c.IndentedJSON(code, gin.H{"msg" : client})
	log.Printf(server + ": %v", err)
}