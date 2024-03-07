package eventhandlers

import (
	"encoding/json"

	"github.com/rodsil01/solucion-apuesta-total/contracts"
	"github.com/rodsil01/solucion-apuesta-total/messagebroker"
	"github.com/rodsil01/solucion-apuesta-total/reservations/domain/service"
)

type BaggageReservationFailedHandler struct {
	service      service.ReservationService
	eventManager messagebroker.EventManager
}

func (h *BaggageReservationFailedHandler) HandleEvent(eventType string, event interface{}) {
	jsonData, _ := json.Marshal(event)

	var payload contracts.BaggageReservationFailed
	marshalErr := json.Unmarshal(jsonData, &payload)

	if marshalErr != nil {
		return
	}

	h.service.DeleteReservation(payload.ReservationId)
}

func NewBaggageReservationFailedHandler(service service.ReservationService, eventManager messagebroker.EventManager) *BaggageReservationFailedHandler {
	return &BaggageReservationFailedHandler{service: service, eventManager: eventManager}
}
