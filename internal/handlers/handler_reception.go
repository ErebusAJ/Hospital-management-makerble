package handlers

import (
	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// registers receptionist
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


func (cfg *apiConfig) updateReceptionist(c *gin.Context) {
	var reqDetails struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, err)
		return
	}
	userID := tempID.(uuid.UUID)

	user, err := cfg.DB.GetReceptionistByID(c, userID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	if reqDetails.Name == "" {
		reqDetails.Name = user.Name
	}
	if reqDetails.Email == "" {
		reqDetails.Email = user.Email
	}
	if reqDetails.Phone == "" {
		reqDetails.Phone = user.Phone
	}
	if reqDetails.Address == "" {
		reqDetails.Address = user.Address
	}

	err = cfg.DB.UpdateReceptionist(c, db.UpdateReceptionistParams{
		Name: reqDetails.Name,
		Email: reqDetails.Email,
		Phone: reqDetails.Phone,
		Address: reqDetails.Address,
		ID: userID,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}
}


func(cfg *apiConfig) getReceptionists(c *gin.Context){
	list, err := cfg.DB.ListReceptionists(c)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(200, list)
}

func(cfg *apiConfig) deleteReceptionist(c *gin.Context){
	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	userID := tempID.(uuid.UUID)

	err := cfg.DB.DeleteReceptionist(c, userID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(204, gin.H{"msg" : "success"})
}