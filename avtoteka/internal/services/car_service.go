package services

import (
	"avtoteka/avtoteka/internal/models"
	"avtoteka/avtoteka/internal/repository"
	"context"
)

// в services находится бизнес логика
// прослойка между обработчиками (HTTP) и репозиторием (БД)

// dependence on the repo
type CarService struct {
	repo *repository.CarRepository // инъекция репозитория
}

// constructor for CarService
func NewCarService(repo *repository.CarRepository) *CarService {
	return &CarService{repo: repo} // initialization
}

func (s *CarService) CreateCar(ctx context.Context, car *models.Car) error {
	return s.repo.CreateCar(ctx, car)
}

func (s *CarService) GetCar(ctx context.Context, id int) (*models.Car, error) {
	return s.repo.GetCarById(ctx, id)
}
