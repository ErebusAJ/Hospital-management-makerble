package handlers

import (
	"log"

	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// logs in user
// returns jwt
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

func(cfg *apiConfig) getUserDetails(c *gin.Context){
	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	userID := tempID.(uuid.UUID)

	userRole, _ := c.Get("userRole")

	var err error
	switch userRole{
	case "receptionist":
		user, err := cfg.DB.GetReceptionistByID(c, userID)
		if err == nil {
			c.IndentedJSON(200, user)
			return
		}
	case "doctor":
		user, err := cfg.DB.GetDoctorByID(c, userID)
		if err == nil {
			c.IndentedJSON(200, user)
			return
		}
	case "patient":
		user, err := cfg.DB.GetPatientByID(c, userID)
		if err == nil {
			c.IndentedJSON(200, user)
			return
		}
	}

	utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
}