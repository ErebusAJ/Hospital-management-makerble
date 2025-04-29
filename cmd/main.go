package main

import (
	"log"
	"os"

	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	// connect to database
	db, err := utils.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// initlaize gin router
	r := gin.Default()

	port := os.Getenv("PORT_NO")
	if port == "" {
		log.Fatal("error getting port no")
	}

	r.Run(":"+port)
}