package models

import "time"

// Coordinates -.
type Coordinates struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}

// Sensor - .
type Sensor struct {
	Index          int         `json:"id,omitempty"`
	Group          string      `json:"group,omitempty"`
	Coordinates    Coordinates `json:"coordinates,omitempty"`
	DataOutputRate int         `json:"data_output_rate,omitempty"`
}

// Fish - .
type Fish struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// DetectedFishes - .
type DetectedFishes struct {
	ID          int    `json:"id,omitempty"`
	SensorIndex int    `json:"sensor_index,omitempty"`
	SensorGroup string `json:"sensor_group,omitempty"`
	FishID      int    `json:"fish_id,omitempty"`
	Count       int    `json:"count,omitempty"`
}

// SensorsDetection - .
type SensorsDetection struct {
	ID               int       `json:"id,omitempty"`
	SensorIndex      int       `json:"sensor_id,omitempty"`
	SensorGroup      string    `json:"sensor_group,omitempty"`
	DetectedFishesID int       `json:"detected_fishes_id,omitempty"`
	Temperature      float64   `json:"temperature,omitempty"`
	Transparency     int       `json:"transparency,omitempty"`
	DetectedAt       time.Time `json:"detection_time,omitempty"`
}
