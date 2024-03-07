package handlerconfig

import (
	"net/http"

	"github.com/gorilla/mux"
	ctx "github.com/rodsil01/solucion-apuesta-total/baggage/domain/context"
	eventhandlers "github.com/rodsil01/solucion-apuesta-total/baggage/event-handlers"
	handlers "github.com/rodsil01/solucion-apuesta-total/baggage/http-handlers"
	"github.com/rodsil01/solucion-apuesta-total/baggage/repository"
	"github.com/rodsil01/solucion-apuesta-total/baggage/service"
	"github.com/rodsil01/solucion-apuesta-total/messagebroker"
)

func (config *HandlerConfig) AddBaggageHandlers(router *mux.Router) {
	context := ctx.GetDbContext()
	eventManager := messagebroker.GetEventManager()

	baggageService := service.NewBaggageServiceImpl(repository.NewBaggageRepositoryImpl(context))

	clientHandler := handlers.NewBaggageHandler(baggageService)

	router.HandleFunc("/baggage/reservations/{reservation_id}", clientHandler.GetBaggageByReservationId).Methods(http.MethodGet)

	eventManager.InjectHandlers(
		eventhandlers.NewReservationCreatedHandler(baggageService, eventManager),
	)
}
