package service

import (
	"github.com/rodsil01/solucion-apuesta-total/baggage/domain/model"
	"github.com/rodsil01/solucion-apuesta-total/baggage/exceptions"
	"github.com/rodsil01/solucion-apuesta-total/contracts"
)

type BaggageService interface {
	GetBaggageByReservationId(string) ([]contracts.BaggageDto, *exceptions.AppError)
	CreateBaggage(newBaggageRequest *contracts.NewBaggageRequest) (*model.Baggage, *exceptions.AppError)
	CreateBaggageRange(newBaggageRequest []*contracts.NewBaggageRequest) ([]string, *exceptions.AppError)
	DeleteBaggage(baggageId string) *exceptions.AppError
	DeleteBaggageRange(baggageIds []string) *exceptions.AppError
}
