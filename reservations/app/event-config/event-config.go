package eventconfig

import (
	"log"

	"github.com/rodsil01/solucion-apuesta-total/messagebroker"
)

func (config *EventConfig) AddRabbitMQ(url string) {
	em := messagebroker.GetEventManager()
	err := em.Initialize(url)

	if err != nil {
		log.Fatal("Error initializing event manager: " + err.Error())
	}

	err = em.ConsumeEvents()

	if err != nil {
		log.Fatal("Error starting events consumer: " + err.Error())
	}
}
