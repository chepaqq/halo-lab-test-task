package group

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chepaqq99/halo-lab-test-task/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type GroupDB struct {
	db    *sqlx.DB
	cache *redis.Client
}

func NewGroupDB(db *sqlx.DB, cache *redis.Client) *GroupDB {
	return &GroupDB{db: db, cache: cache}
}

// GetAverageTransparency - .
func (r *GroupDB) GetAverageTransparency(groupName string) (float64, error) {
	// Check if the result is cached
	cachedResult, err := r.cache.Get(context.Background(), fmt.Sprintf("average_transparency:%s", groupName)).Result()
	if err == nil {
		// Result found in cache, return the cached value
		averageTransparency, err := strconv.ParseFloat(cachedResult, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse cached result: %s", err)
		}
		return averageTransparency, nil
	}

	// Result not found in cache, fetch from the database
	var averageTransparency float64
	query := `SELECT AVG(transparency) FROM sensor
		JOIN sensor_group g ON g.id = group_id WHERE g.name = $1`
	if err := r.db.Get(&averageTransparency, query, groupName); err != nil {
		return 0, err
	}

	// Cache the result
	err = r.cache.Set(context.Background(), fmt.Sprintf("average_transparency:%s", groupName), averageTransparency, 10*time.Second).Err()
	if err != nil {
		return 0, fmt.Errorf("failed to cache result: %s", err)
	}

	return averageTransparency, nil
}

// GetAverageTemperature - .
func (r *GroupDB) GetAverageTemperature(groupName string) (float64, error) {
	// Check if the result is cached
	cachedResult, err := r.cache.Get(context.Background(), fmt.Sprintf("average_temperature:%s", groupName)).Result()
	if err == nil {
		// Result found in cache, return the cached value
		averageTemperature, err := strconv.ParseFloat(cachedResult, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse cached result: %w", err)
		}
		return averageTemperature, nil
	}

	// Result not found in cache, fetch from the database
	var averageTemperature float64
	query := `SELECT AVG(temperature) FROM sensor
		JOIN sensor_group g ON g.id = group_id WHERE g.name = $1`
	if err = r.db.Get(&averageTemperature, query, groupName); err != nil {
		return 0, err
	}

	// Cache the result
	err = r.cache.Set(context.Background(), fmt.Sprintf("average_temperature:%s", groupName), averageTemperature, 10*time.Second).Err()
	if err != nil {
		return 0, fmt.Errorf("failed to cache result: %w", err)
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

// GetListOfSpecies - .
func (r *GroupDB) GetListOfSpecies(groupName string) (map[string]int, error) {
	detectedFishes := make(map[string]int)
	query := `SELECT fish.name, count FROM detected_fish
			JOIN fish ON fish_id = fish.id
			JOIN sensor_group ON sensor_group.id = sensor_group_id
			WHERE sensor_group.name = $1`
	rows, err := r.db.Queryx(query, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fishName string
		var count int

		err := rows.Scan(&fishName, &count)
		if err != nil {
			return nil, err
		}

		detectedFishes[fishName] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return detectedFishes, nil
}

// GetTopListOfSpecies - .
func (r *GroupDB) GetTopListOfSpecies(groupName string, top int) (map[string]int, error) {
	detectedFishes := make(map[string]int)
	query := `SELECT fish.name, count FROM detected_fish
			JOIN fish ON fish_id = fish.id
			JOIN sensor_group ON sensor_group.id = sensor_group_id
			WHERE sensor_group.name = $1
			ORDER BY detected_fish.count DESC
			LIMIT $2`
	rows, err := r.db.Queryx(query, groupName, top)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fishName string
		var count int

		err := rows.Scan(&fishName, &count)
		if err != nil {
			return nil, err
		}

		detectedFishes[fishName] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return detectedFishes, nil
}

// Create - .
func (r *GroupDB) Create(groupName string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	query := `INSERT INTO sensor_group(name) VALUES ($1)`
	row := tx.QueryRow(query, groupName)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	return id, tx.Commit()
}

// GetLastIDInGroup - .
func (r *GroupDB) GetLastIDInGroup(groupID int) (int, error) {
	var lastIndex int
	query := `SELECT index FROM sensor WHERE group_id = $1 ORDER BY index DESC LIMIT 1`
	if err := r.db.Get(&lastIndex, query, groupID); err != nil {
		return 0, err
	}
	return lastIndex, nil
}
