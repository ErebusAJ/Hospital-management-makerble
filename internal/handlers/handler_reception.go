package handlers

import (
	"log"

	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func (cfg *apiConfig) login(c *gin.Context) {
	var reqDetails struct {
		Email string `json:"email" binding:"required"`
		Pass  string `json:"password" binding:"required"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	if receptionist, err := cfg.DB.GetReceptionistByEmail(c, reqDetails.Email); err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(receptionist.PasswordHash), []byte(reqDetails.Pass)); err == nil {
			// Generate token & respond
			token, e := utils.GenerateJWT(receptionist.ID, "receptionist")
			log.Printf("error %v", e)
			c.IndentedJSON(200, gin.H{"token": token})
			return
		}
	}

	// Try doctor
	if doctor, err := cfg.DB.GetDoctorByEmail(c, reqDetails.Email); err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(doctor.PasswordHash), []byte(reqDetails.Pass)); err == nil {
			// Generate token & respond
			token, e := utils.GenerateJWT(doctor.ID, "doctor")
			log.Printf("error %v", e)
			c.IndentedJSON(200, gin.H{"token": token})
			return
		}
	}

	// Try patient
	if patient, err := cfg.DB.GetPatientByEmail(c, reqDetails.Email); err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(patient.PasswordHash), []byte(reqDetails.Pass)); err == nil {
			// Generate token & respond
			token, e := utils.GenerateJWT(patient.ID, "patient")
			log.Printf("error %v", e)
			c.IndentedJSON(200, gin.H{"token": token})
			return
		}
	}

	utils.ErrorJSON(c, 401, utils.UnauthorizedError, "invalid email or password", nil)
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
}
