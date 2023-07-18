package group

import (
	"github.com/jmoiron/sqlx"
)

type GroupDB struct {
	db *sqlx.DB
}

func NewGroupDB(db *sqlx.DB) *GroupDB {
	return &GroupDB{db: db}
}

// GetAverageTransparency - .
func (r *GroupDB) GetAverageTransparency(groupName string) (float64, error) {
	var averageTransparency float64
	query := "SELECT AVG(transparency) FROM sensor_detection WHERE sensor_group = $1"
	if err := r.db.Get(&averageTransparency, query, groupName); err != nil {
		return -1, err
	}

	return averageTransparency, nil
}

// GetAverageTemperature - .
func (r *GroupDB) GetAverageTemperature(groupName string) (float64, error) {
	var averageTemperature float64
	query := "SELECT AVG(temperature) FROM sensor_detection WHERE sensor_group = $1"
	if err := r.db.Get(&averageTemperature, query, groupName); err != nil {
		return -1, err
	}

	return averageTemperature, nil
}
