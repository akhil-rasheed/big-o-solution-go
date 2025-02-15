package main

import (
	"errors"
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	router.GET("/health", healthCheck)
	router.GET("/:location_id", getRecord)
	router.PUT("/:location_id", putRecord)

	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getRecord(c *gin.Context) {
	locationID := c.Param("location_id")
	store.mu.RLock()
	data, exists := store.data[locationID]
	store.mu.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func putRecord(c *gin.Context) {
	locationID := c.Param("location_id")
	var input SensorDataInput

	if err := updateRecord(locationID, input); err != nil {
		switch {
		case errors.Is(err, ErrInsufficientStorage):
			c.JSON(http.StatusInsufficientStorage, gin.H{"error": "Insufficient storage"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Write rejected"})
		}
		return
	}

	c.Status(http.StatusCreated)
}

var ErrInsufficientStorage = errors.New("insufficient storage")

func updateRecord(locationID string, input SensorDataInput) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if len(store.data) >= maxStorageSize && store.data[locationID].ID == uuid.Nil {
		return ErrInsufficientStorage
	}

	record, exists := store.data[locationID]
	if !exists {
		record = SensorData{
			LocationID:        locationID,
			ModificationCount: 0,
		}
	}

	record.ID = input.ID
	record.SeismicActivity = input.SeismicActivity
	record.TemperatureC = input.TemperatureC
	record.RadiationLevel = input.RadiationLevel
	record.ModificationCount++

	store.data[locationID] = record

	return nil
}
