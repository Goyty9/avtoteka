package repository

import (
	"avtoteka/avtoteka/internal/models"
	"context"
	"database/sql"
	"errors"
)

// all work with DB

// structure for work with Car
// db connection
type CarRepository struct {
	db *sql.DB
}

// constructor for CarRepository
// return repository instance
func NewCarRepository(db *sql.DB) *CarRepository {
	return &CarRepository{db: db}
}

// create new car in DB
func (r *CarRepository) CreateCar(ctx context.Context, car *models.Car) error {
	query := `
			INSERT INTO cars (
			        manufacturer_id, brand, model, created_year,
			        engine_capacity, power, drive_type, transmission_type
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING id
			`

	return r.db.QueryRowContext(ctx, query,
		car.ManufacturerID, car.Brand, car.Model, car.CreatedYear,
		car.EngineCapacity, car.Power, car.DriveType, car.TransmissionType,
	).Scan(&car.ID)
}

func (r *CarRepository) GetCarById(ctx context.Context, id int) (*models.Car, error) {
	query := `
			SELECT
					id, manufacturer_id, brand, model, created_year,
			        engine_capacity, power, drive_type, transmission_type, created_at, updated_at
			FROM cars
			WHERE id = $1
			`

	car := &models.Car{} // empty struct for save data from DB
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&car.ID,
		&car.ManufacturerID,
		&car.Brand,
		&car.Model,
		&car.CreatedYear,
		&car.EngineCapacity,
		&car.Power,
		&car.DriveType,
		&car.TransmissionType,
		&car.CreatedAt,
		&car.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return car, err
}
