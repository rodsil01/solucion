package service

import (
	"github.com/rodsil01/solucion-apuesta-total/reservations/domain/model"
	"github.com/rodsil01/solucion-apuesta-total/reservations/exceptions"
)

type RouteService interface {
	GetAllRoutes() ([]model.Route, *exceptions.AppError)
	GetRouteById(string) (*model.Route, *exceptions.AppError)
}
