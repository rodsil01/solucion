package repository

import (
	"github.com/rodsil01/solucion-apuesta-total/reservations/domain/model"
	"github.com/rodsil01/solucion-apuesta-total/reservations/exceptions"
)

type ReservationRepository interface {
	FindAll() ([]model.Reservation, *exceptions.AppError)
	FindById(string) (*model.Reservation, *exceptions.AppError)
	Create(reservation model.Reservation) (*model.Reservation, *exceptions.AppError)
	Delete(reservationId string) *exceptions.AppError
}
