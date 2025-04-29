package handlers

import (
	"os"

	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/middleware"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *db.Queries
}

func RegisterRoutes(r *gin.Engine) {
	godotenv.Load()
	
	DB, _ := utils.ConnectDB()
	
	dbQueries := db.New(DB)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	r.POST("/v1/reception/register", apiCfg.registerReceptionist)
	r.POST("/v1/login", apiCfg.login)

	protected := r.Group("/v1/auth")

	sign := os.Getenv("SECRET_KEY")
	protected.Use(middleware.AuthMiddleware(sign))
	{
		// global
		protected.GET("/user-details", apiCfg.getUserDetails)

		// receptionist routes
		protected.PUT("/receptionist", apiCfg.updateReceptionist)
	}

}