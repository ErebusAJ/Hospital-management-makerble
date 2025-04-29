package handlers

import (
	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func (cfg *apiConfig) registerReceptionist(c *gin.Context) {
	var reqDetails struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address string `json:"address" binding:"required"`
		Pass    string `json:"password" binding:"required"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	hashedPass, _ := utils.PasswordHash(reqDetails.Pass)

	err = cfg.DB.CreateReceptionist(c, db.CreateReceptionistParams{
		Name:         reqDetails.Name,
		Email:        reqDetails.Email,
		Phone:        reqDetails.Phone,
		Address:      reqDetails.Address,
		PasswordHash: hashedPass,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(201, gin.H{"message": "success"})
}
