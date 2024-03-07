package eventhandlers

import (
	"encoding/json"

	"github.com/rodsil01/solucion-apuesta-total/baggage/domain/service"
	"github.com/rodsil01/solucion-apuesta-total/contracts"
	"github.com/rodsil01/solucion-apuesta-total/messagebroker"
)

type ReservationCreatedHandler struct {
	service      service.BaggageService
	eventManager messagebroker.EventManager
}

func (h *ReservationCreatedHandler) HandleEvent(eventType string, event interface{}) {
	jsonData, _ := json.Marshal(event)

	var payload contracts.ReservationCreated
	marshalErr := json.Unmarshal(jsonData, &payload)

	if marshalErr != nil {
		return
	}

	var baggage []*contracts.NewBaggageRequest

	for _, r := range payload.Baggage {
		b := &contracts.NewBaggageRequest{
			ReservationId: payload.ReservationId,
			Description:   r.Description,
			Weight:        r.Weight,
		}
		baggage = append(baggage, b)
	}

	ids, err := h.service.CreateBaggageRange(baggage)

	if err != nil {
		h.eventManager.PublishEvent(contracts.BAGGAGE_RESERVATION_FAILED, contracts.BaggageReservationFailed{
			Ids:           ids,
			ReservationId: payload.ReservationId,
		})
	}
}

func NewReservationCreatedHandler(service service.BaggageService, eventManager messagebroker.EventManager) *ReservationCreatedHandler {
	return &ReservationCreatedHandler{service: service, eventManager: eventManager}
}
