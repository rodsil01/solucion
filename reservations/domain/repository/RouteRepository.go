package repository

import (
	"github.com/rodsil01/solucion-apuesta-total/reservations/domain/model"
	"github.com/rodsil01/solucion-apuesta-total/reservations/exceptions"
)

type RouteRepository interface {
	FindAll() ([]model.Route, *exceptions.AppError)
	FindById(string) (*model.Route, *exceptions.AppError)
}
