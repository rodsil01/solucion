package repository

import (
	"github.com/rodsil01/solucion-apuesta-total/reservations/domain/model"
	"github.com/rodsil01/solucion-apuesta-total/reservations/exceptions"
)

type ClientRepository interface {
	FindAll() ([]model.Client, *exceptions.AppError)
	FindById(string) (*model.Client, *exceptions.AppError)
}
