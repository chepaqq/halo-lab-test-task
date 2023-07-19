package region

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RegionDB struct {
	db *sqlx.DB
}

func NewRegionDB(db *sqlx.DB) *RegionDB {
	return &RegionDB{db: db}
}

// GetMinTemperatureInRegion - .
func (r *RegionDB) GetMinTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error) {
	var minTemperature float64
	fmt.Print(xMin, yMin, zMin, xMax, yMax, zMax)
	query := `SELECT MIN(temperature) FROM sensor JOIN coordinates
		ON sensor.coordinates_id = coordinates.id
		WHERE coordinates.x >= $1 AND coordinates.x <= $2
		AND coordinates.y >= $3 AND coordinates.y <= $4
		AND coordinates.z >= $5 AND coordinates.z <= $6`

	fmt.Print(query)
	if err := r.db.Get(&minTemperature, query, xMin, xMax, yMin, yMax, zMin, zMax); err != nil {
		return 0, err
	}

	return minTemperature, nil
}

// GetMaxTemperatureInRegion - .
func (r *RegionDB) GetMaxTemperatureInRegion(xMin, yMin, zMin, xMax, yMax, zMax int) (float64, error) {
	var maxTemperature float64
	query := `SELECT MAX(temperature) FROM sensor JOIN coordinates
		ON sensor.coordinates_id = coordinates.id
		WHERE coordinates.x >= $1 AND coordinates.x <= $2
		AND coordinates.y >= $3 AND coordinates.y <= $4
		AND coordinates.z >= $5 AND coordinates.z <= $6`
	if err := r.db.Get(&maxTemperature, query, xMin, xMax, yMin, yMax, zMin, zMax); err != nil {
		return 0, err
	}

	return maxTemperature, nil
}
