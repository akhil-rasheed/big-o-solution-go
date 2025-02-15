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

type InMemoryStore struct {
	data map[string]SensorData
	mu   sync.RWMutex
}

var store = &InMemoryStore{
	data: make(map[string]SensorData),
}

const (
	maxStorageSize      = 1000 
)

func main() {
	router := gin.Default()
	router.GET("/records", getRecord)
	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getRecord(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sensorData)
}