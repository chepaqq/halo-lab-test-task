package group

import (
	"github.com/chepaqq99/halo-lab-test-task/internal/models"
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
	query := `SELECT AVG(transparency) FROM sensor
		JOIN sensor_group g ON g.id = group_id WHERE g.name = $1`
	if err := r.db.Get(&averageTransparency, query, groupName); err != nil {
		return 0, err
	}

	return averageTransparency, nil
}

// GetAverageTemperature - .
func (r *GroupDB) GetAverageTemperature(groupName string) (float64, error) {
	var averageTemperature float64
	query := `SELECT AVG(temperature) FROM sensor
		JOIN sensor_group g ON g.id = group_id WHERE g.name = $1`
	if err := r.db.Get(&averageTemperature, query, groupName); err != nil {
		return 0, err
	}

	return averageTemperature, nil
}

// GetGroupByName - .
func (r *GroupDB) GetGroupByName(groupName string) (*models.SensorGroup, error) {
	var group models.SensorGroup
	query := `SELECT * FROM sensor_group WHERE name = $1`
	if err := r.db.Get(&group, query, groupName); err != nil {
		return &group, err
	}
	return &group, nil
}
