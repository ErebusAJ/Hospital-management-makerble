package handlers

import (
	"sync"

	"github.com/ErebusAJ/makerble-backend/internal/db"
	"github.com/ErebusAJ/makerble-backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// functions to handle patients

// register patients
func(cfg *apiConfig) registerPatient(c *gin.Context) {
	var reqDetails struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address	string `json:"address" binding:"required"`
		DocID	string `json:"doctor_id" binding:"required"`
		Pass    string `json:"password" binding:"required"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var receptionistID uuid.UUID
	var doctorID uuid.UUID
	var errR bool
	var errD error
	// get user id from token(middleware)
	go func() {
		wg.Done()
		var tempID any
		tempID, errR = c.Get("userID")
		receptionistID = tempID.(uuid.UUID)
	}()
	// parse doctor id
	go func() {
		wg.Done()
		doctorID, errD = uuid.Parse(reqDetails.DocID)
	}()
	// hash password
	hashedPass, _ := utils.PasswordHash(reqDetails.Pass)
	

	wg.Wait()

	if !errR || errD != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	err = cfg.DB.RegisterPatient(c, db.RegisterPatientParams{
		Name: reqDetails.Name,
		Email: reqDetails.Email,
		Phone: reqDetails.Phone,
		Address: reqDetails.Address,
		ReceptionistID: receptionistID,
		DoctorID: doctorID,
		PasswordHash: hashedPass,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(201, gin.H{"message" : "success"})
}

// update existing patient details
func(cfg *apiConfig) updatePatient(c* gin.Context) {
	var reqDetails struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Phone   string `json:"phone" binding:"required"`
		Address	string `json:"address" binding:"required"`
		DocID	string `json:"doctor_id" binding:"required"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	tempID := c.Param("patientID")
	patientID, err := uuid.Parse(tempID)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	user, err := cfg.DB.GetPatientByID(c, patientID)
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
	if reqDetails.Address == "" {
		reqDetails.Address = user.Address
	}
	var doctorID uuid.UUID
	if reqDetails.DocID == "" {
		doctorID = user.DoctorID
	}else {
		doctorID, err = uuid.Parse(reqDetails.DocID)
		if err != nil {
			utils.ErrorJSON(c, 400, utils.InvalidError, "parsing error", err)
			return
		}
	}

	err = cfg.DB.UpdatePatient(c, db.UpdatePatientParams{
		Name: reqDetails.Name,
		Email: reqDetails.Email,
		Phone: reqDetails.Phone,
		Address: reqDetails.Address,
		DoctorID: doctorID,
		ID: patientID,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(203, gin.H{"message" : "success"})
}

// get all patients 
func(cfg *apiConfig) getPatientList(c *gin.Context) {
	patients, err := cfg.DB.ListPatients(c)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(200, patients)	
}

// get patient by id
func(cfg *apiConfig) getPatientByID(c *gin.Context) {
	tempID := c.Param("patientID")
	patientID, err := uuid.Parse(tempID)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	patient, err := cfg.DB.GetPatientByID(c, patientID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(200, patient)	
}

// delete patient
func(cfg *apiConfig) deletePatient(c *gin.Context) {
	tempID := c.Param("patientID")
	patientID, err := uuid.Parse(tempID)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	err = cfg.DB.DeletePatient(c, patientID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, "hashing error", err)
		return
	}

	c.IndentedJSON(204, gin.H{"message" : "success"})	
}

