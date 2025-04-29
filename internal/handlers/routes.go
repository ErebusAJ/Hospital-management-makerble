package handlers

import (
	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type apiConfig struct {
	DB *db.Queries
}

func RegisterRoutes(r *gin.Engine) {
	
	DB, _ := utils.ConnectDB()
	
	dbQueries := db.New(DB)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	r.POST("/v1/reception/register", apiCfg.registerReceptionist)

}