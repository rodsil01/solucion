module github.com/rodsil01/solucion-apuesta-total/baggage

go 1.22.1

require (
	github.com/rodsil01/solucion-apuesta-total/contracts v1.0.0
	github.com/rodsil01/solucion-apuesta-total/messagebroker v1.0.0
)

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/streadway/amqp v1.1.0 // indirect
)

replace (
	github.com/rodsil01/solucion-apuesta-total/contracts => ../contracts
	github.com/rodsil01/solucion-apuesta-total/messagebroker => ../message-broker
)
