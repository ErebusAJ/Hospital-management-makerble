package handlers

import (
	"database/sql"
	"log"

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

// get patients assigned to particular doctor
func(cfg *apiConfig) getPatientsByDoctor(c *gin.Context) {
	tempID := c.Param("doctorID")
	doctorID, err := uuid.Parse(tempID)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	patients, err :=	cfg.DB.GetPatientsByDoctor(c, doctorID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(200, patients)
}

// register user medical data
func(cfg *apiConfig) registerMedicalData(c *gin.Context) {
	var reqDetails struct {
		Symptoms	string `json:"symptoms" binding:"required"`
		Diagnosis	string `json:"diagnosis" binding:"required"`
		Prescription   string `json:"prescription" binding:"required"`
		Notes	string `json:"notes" binding:"required"`
		Tests	string `json:"tests" binding:"required"`
		FollowUp    string `json:"follow_up_date" binding:"required"`
	}

	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	doctorID := tempID.(uuid.UUID)

	tempID = c.Param("patientID")
	patientID, err := uuid.Parse(tempID.(string))
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	_, err = cfg.DB.GetDoctorByID(c, doctorID)
	if err != nil { 
		if err == sql.ErrNoRows {
			utils.ErrorJSON(c, 401, utils.UnauthorizedError, "no doctor found with specified id", err)
		} else {
			utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		}
		return
	}

	err = c.BindJSON(&reqDetails)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	// handle sql nullstring
	notes := sql.NullString{
		String: reqDetails.Notes,
		Valid: reqDetails.Notes != "",
	}

	err = cfg.DB.CreatePatientHistory(c, db.CreatePatientHistoryParams{
		PatientID: patientID,
		DoctorID: doctorID,
		Symptoms: reqDetails.Symptoms,
		Diagnosis: reqDetails.Diagnosis,
		Prescription: reqDetails.Prescription,
		Notes: notes,
		TestsRecommended: reqDetails.Tests,
		FollowUpDate: reqDetails.FollowUp,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(200, gin.H{"message" : "success"})
}


func(cfg *apiConfig) updatePatientMedicalData(c *gin.Context) {
	var reqDetails struct {
		Symptoms	string `json:"symptoms"`
		Diagnosis	string `json:"diagnosis"`
		Prescription   string `json:"prescription"`
		Notes	string `json:"notes"`
		Tests	string `json:"tests"`
		FollowUp    string `json:"follow_up_date"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		log.Printf("its bind")
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	// verify user sending request is doctor
	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	doctorID := tempID.(uuid.UUID)

	_, err = cfg.DB.GetDoctorByID(c, doctorID)
	if err != nil { 
		if err == sql.ErrNoRows {
			utils.ErrorJSON(c, 401, utils.UnauthorizedError, "no doctor found with specified id", err)
		} else {
			utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		}
		return
	}

	tempMID := c.Param("medicalID")
	medicalID, err := uuid.Parse(tempMID)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	patientHistory, err := cfg.DB.GetPatientHistoryByID(c, medicalID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	if reqDetails.Symptoms == "" {
		reqDetails.Symptoms = patientHistory.Symptoms
	}
	if reqDetails.Diagnosis == "" {
		reqDetails.Diagnosis = patientHistory.Diagnosis
	}
	if reqDetails.Prescription == "" {
		reqDetails.Prescription = patientHistory.Prescription
	}
	var notes sql.NullString
	if reqDetails.Notes == "" {
		notes = sql.NullString{
			String: reqDetails.Notes,
			Valid: reqDetails.Notes != "",
		}
	}else{ 
		notes = sql.NullString{
			String: patientHistory.Notes.String,
			Valid: reqDetails.Notes != "",
		}
	}
	if reqDetails.Notes == "" {
		reqDetails.Notes = patientHistory.Notes.String
	}
	if reqDetails.Tests == "" {
		reqDetails.Tests = patientHistory.TestsRecommended
	}
	if reqDetails.FollowUp == "" {
		reqDetails.FollowUp = patientHistory.FollowUpDate
	}

	err = cfg.DB.UpdatePatientHistory(c, db.UpdatePatientHistoryParams{
		Symptoms: reqDetails.Symptoms,
		Diagnosis: reqDetails.Diagnosis,
		Prescription: reqDetails.Prescription,
		FollowUpDate: reqDetails.FollowUp,
		Notes: notes,
		TestsRecommended: reqDetails.Tests,
		ID: medicalID,
	})
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}
	
	c.IndentedJSON(204, gin.H{"message" : "success"})
}


// get patiend medical data
func(cfg *apiConfig) getPatientMedicalData(c *gin.Context) {
	tempID := c.Param("patientID")
	patientID, err := uuid.Parse(tempID)
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	pateintData, err := cfg.DB.ListPatientHistoryByPatientID(c, patientID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(200, pateintData)
}


func(cfg *apiConfig) deletePatientMedicalData(c *gin.Context) {
	// verify user is doctor
	tempID, exists := c.Get("userID")
	if !exists {
		utils.ErrorJSON(c, 401, utils.UnauthorizedError, utils.MiddlewareError, nil)
		return
	}
	doctorID := tempID.(uuid.UUID)

	_, err := cfg.DB.GetDoctorByID(c, doctorID)
	if err != nil { 
		if err == sql.ErrNoRows {
			utils.ErrorJSON(c, 401, utils.UnauthorizedError, "no doctor found with specified id", err)
		} else {
			utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		}
		return
	}

	tempID = c.Param("medicalID")
	medicalID, err := uuid.Parse(tempID.(string))
	if err != nil {
		utils.ErrorJSON(c, 400, utils.InvalidError, utils.RequestBodyError, err)
		return
	}

	err = cfg.DB.DeletePatientHistory(c, medicalID)
	if err != nil {
		utils.ErrorJSON(c, 500, utils.InternalError, utils.DatabaseError, err)
		return
	}

	c.IndentedJSON(204, gin.H{"message" : "success"})
}