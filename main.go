package main

import (
	"github.com/google/uuid"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SensorDataFields struct {
	ID               uuid.UUID `json:"id"`
	ModificationCount int      `json:"modification_count"`
	SeismicActivity   float64  `json:"seismic_activity"`
	TemperatureC      float64  `json:"temperature_c"`
	RadiationLevel    float64  `json:"radiation_level"`
}

var sensorData = []SensorDataFields{
	{
		ID:               uuid.Must(uuid.Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")),
		ModificationCount: 4,
		SeismicActivity:   23.4,
		TemperatureC:      32.6,
		RadiationLevel:    3.22,
	},
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecord)
	router.Run("localhost:8080")
}

func getRecord(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sensorData)
}