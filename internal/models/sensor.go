package models

import "time"

// Coordinates -.
type Coordinates struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}

// SensorGroup - .
type SensorGroup struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Sensor - .
type Sensor struct {
	Index          int           `json:"index,omitempty"`
	GroupID        int           `json:"group_id,omitempty"`
	Coordinates    Coordinates   `json:"coordinates,omitempty"`
	DataOutputRate time.Duration `json:"data_output_rate,omitempty"`
	Temperature    float64       `json:"temperature,omitempty"`
	Transparency   int           `json:"transparency,omitempty"`
}

// Fish - .
type Fish struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// DetectedFishes - .
type DetectedFishes struct {
	ID            int `json:"id,omitempty"`
	FishID        int `json:"fish_id,omitempty"`
	Count         int `json:"count,omitempty"`
	SensorIndex   int `json:"sensor_index,omitempty"`
	SensorGroupID int `json:"sensor_group_id,omitempty"`
}
