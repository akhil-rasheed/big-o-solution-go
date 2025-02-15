package main

import (
	"github.com/google/uuid"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SensorData struct {
	ID                uuid.UUID `json:"id"`
	ModificationCount int       `json:"modification_count"`
	SeismicActivity   float64   `json:"seismic_activity"`
	TemperatureC      float64   `json:"temperature_c"`
	RadiationLevel    float64   `json:"radiation_level"`
	LocationID        string    `json:"location_id"`
}

type SensorDataInput struct {
	ID              uuid.UUID `json:"id"`
	SeismicActivity float64   `json:"seismic_activity"`
	TemperatureC    float64   `json:"temperature_c"`
	RadiationLevel  float64   `json:"radiation_level"`
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecord)
	router.Run("localhost:8080")
}

func getRecord(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sensorData)
}