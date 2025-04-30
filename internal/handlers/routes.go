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

	r.POST("/v1/register/reception", apiCfg.registerReceptionist)
	r.POST("/v1/register/doctor", apiCfg.registerDoctor) 
	r.POST("/v1/login", apiCfg.login)

	protected := r.Group("/v1/auth")

	sign := os.Getenv("SECRET_KEY")
	protected.Use(middleware.AuthMiddleware(sign))
	{
		// global
		protected.GET("/user-details", apiCfg.getUserDetails)

		// receptionist routes
		protected.GET("/receptionist/all", apiCfg.getReceptionistList)
		protected.PUT("/receptionist", apiCfg.updateReceptionist)
		protected.DELETE("/receptionist", apiCfg.deleteReceptionist)

		// doctors routes
		protected.GET("/doctor/all", apiCfg.getDoctorList)
		protected.PUT("/doctor", apiCfg.updateDoctor)
		protected.DELETE("/doctor", apiCfg.deleteDoctor)
		

	}

}