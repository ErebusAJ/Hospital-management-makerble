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

		// doctor routes
		protected.GET("/doctor/all", apiCfg.getDoctorList)
		protected.PUT("/doctor", apiCfg.updateDoctor)
		protected.DELETE("/doctor", apiCfg.deleteDoctor)
		protected.GET("/doctor/:doctorID/patient", apiCfg.getPatientsByDoctor)


		// patient routes
		protected.POST("/patient/register", apiCfg.registerPatient)
		protected.GET("/patient/all", apiCfg.getPatientList)
		protected.GET("/patient/:patientID", apiCfg.getPatientByID)
		protected.PUT("/patient/:patientID", apiCfg.updatePatient)
		protected.DELETE("/patient/:patientID", apiCfg.deletePatient)

		protected.POST("/patient/:patientID/history", apiCfg.registerMedicalData)
		protected.GET("/patient/:patientID/history", apiCfg.getPatientMedicalData)
		protected.PUT("/patient/history/:medicalID", apiCfg.updatePatientMedicalData)
		protected.DELETE("/patient/history/:medicalID", apiCfg.deletePatientMedicalData)
	}

}

func RegisterFrontend(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.StaticFile("/doctor", "./templates/doctor.html")
	r.StaticFile("/receptionist", "./templates/receptionist.html")
	r.StaticFile("/history", "./templates/history.html")
	r.StaticFile("/register", "./templates/register.html")

	r.GET("/login", loginHTML)


}