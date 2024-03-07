package service

import (
	"github.com/rodsil01/solucion-apuesta-total/reservations/domain/model"
	"github.com/rodsil01/solucion-apuesta-total/reservations/exceptions"
)

type ClientService interface {
	GetAllClients() ([]model.Client, *exceptions.AppError)
	GetClientById(string) (*model.Client, *exceptions.AppError)
}
