package sensor

import (
	"github.com/chepaqq99/halo-lab-test-task/internal/models"
	"github.com/jmoiron/sqlx"
)

type SensorDB struct {
	db *sqlx.DB
}

func NewSensorDB(db *sqlx.DB) *SensorDB {
	return &SensorDB{db: db}
}

// CreateSensor - .
func (r *SensorDB) CreateSensor(index, groupID, coordinatesID, dataOutputRate int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := `INSERT INTO sensor(index, group_id, coordinates_id, data_output_rate) VALUES ($1, $2, $3, $4)`
	result, err := tx.Exec(query, index, groupID, coordinatesID, dataOutputRate)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), tx.Commit()
}

// UpdateSensor - .
func (r *SensorDB) UpdateSensor(index, groupID, transparency int, temperature float64) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := `UPDATE sensor SET transparency = $1, temperature = $2 WHERE index = $3 AND group_id = $4`
	result, err := tx.Exec(query, transparency, temperature, index, groupID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), tx.Commit()
}

// CreateCoordinates - .
func (r *SensorDB) CreateCoordinates(coordinates *models.Coordinates) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := `INSERT INTO coordinates(x, y, z) VALUES ($1, $2, $3)`
	result, err := tx.Exec(query, coordinates.X, coordinates.Y, coordinates.Z)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), tx.Commit()
}

func (r *SensorDB) CreateFish(fishName string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := `INSERT INTO fish(name) VALUES ($1)`
	result, err := tx.Exec(query, fishName)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), tx.Commit()
}

func (r *SensorDB) CreateDetectedFishes(fishID, count, sensorIndex, sensorGroupID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := `INSERT INTO detected_fish(fish_id, count, sensor_index, sensor_group_id)
		VALUES ($1, $2, $3, $4)`
	result, err := tx.Exec(query, fishID, count, sensorIndex, sensorGroupID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), tx.Commit()
}

func (r *SensorDB) GetFishBySpecie(fishSpecie string) (*models.Fish, error) {
	var fish models.Fish
	query := `SELECT * FROM fish WHERE name = $1`
	if err := r.db.Get(&fish, query, fishSpecie); err != nil {
		return &fish, err
	}
	return &fish, nil
}
