package handlers

import (
	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// register doctor to database
func(cfg *apiConfig) registerDoctor(c *gin.Context) {
	var reqDetails struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Degree	string `json:"degree" binding:"required"`
		Splz	string `json:"specialization" binding:"required"`
		Pass    string `json:"password" binding:"required"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	hashedPass, err := utils.PasswordHash(reqDetails.Pass)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	err = cfg.DB.RegisterDoctor(c, db.RegisterDoctorParams{
		Name: reqDetails.Name,
		Email: reqDetails.Email,
		Phone: reqDetails.Phone,
		Degree: reqDetails.Degree,
		Specialization: reqDetails.Splz,
		PasswordHash: hashedPass,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(201, gin.H{"message" : "success"})
}

// updates existing doctor details
func(cfg *apiConfig) updateDoctor(c *gin.Context) {
	var reqDetails struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Degree	string `json:"degree"`
		Splz	string `json:"specialization"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	userID := tempID.(uuid.UUID)

	user, err := cfg.DB.GetDoctorByID(c, userID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
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
	if reqDetails.Degree == "" {
		reqDetails.Degree = user.Degree
	}
	if reqDetails.Splz == "" {
		reqDetails.Splz = user.Specialization
	}

	err = cfg.DB.UpdateDoctor(c, db.UpdateDoctorParams{
		Name: reqDetails.Name,
		Email: reqDetails.Email,
		Phone: reqDetails.Phone,
		Degree: reqDetails.Degree,
		Specialization: reqDetails.Splz,
		ID: userID,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}
	
	c.IndentedJSON(204, gin.H{"message" : "success"})
}

// get list of registered doctors
func(cfg *apiConfig) getDoctorList(c *gin.Context){
	doctors, err := cfg.DB.ListDoctors(c)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(200, doctors)
}

// delete doctor
func(cfg *apiConfig) deleteDoctor(c *gin.Context){
	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	userID := tempID.(uuid.UUID)

	err := cfg.DB.DeleteDoctor(c, userID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(204, gin.H{"message" : "success"})
}