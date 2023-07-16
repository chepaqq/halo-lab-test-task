package models

import "time"

type Coordinates struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}

type Sensor struct {
	ID             int         `json:"id,omitempty"`
	Codename       string      `json:"codename,omitempty"`
	Coordinates    Coordinates `json:"coordinates,omitempty"`
	DataOutputRate int         `json:"data_output_rate,omitempty"`
	Group          string      `json:"group,omitempty"`
}

type Fish struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SensorsDetection struct {
	ID             int            `json:"id,omitempty"`
	SensorID       int            `json:"sensor_id,omitempty"`
	Temperature    float64        `json:"temperature,omitempty"`
	Transparency   int            `json:"transparency,omitempty"`
	DetectedFishes map[string]int `json:"detected_fishes,omitempty"`
	DetectionTime  time.Time      `json:"detection_time,omitempty"`
}
