package handlerconfig

import (
	"net/http"

	"github.com/gorilla/mux"
	ctx "github.com/rodsil01/solucion-apuesta-total/reservations/domain/context"
	handlers "github.com/rodsil01/solucion-apuesta-total/reservations/http-handlers"
	"github.com/rodsil01/solucion-apuesta-total/reservations/repository"
	"github.com/rodsil01/solucion-apuesta-total/reservations/service"
)

func (config *HandlerConfig) AddClientHandlers(router *mux.Router) {
	context := ctx.GetDbContext()

	clientHandler := handlers.NewClientHandler(service.NewClientServiceImpl(repository.NewClientRepositoryImpl(context)))

	router.HandleFunc("/clients", clientHandler.GetAllClients).Methods(http.MethodGet)
	router.HandleFunc("/clients/{client_id}", clientHandler.GetClientById).Methods(http.MethodGet)
}
