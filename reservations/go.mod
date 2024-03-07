module github.com/rodsil01/solucion-apuesta-total/reservations

go 1.22.1

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/rodsil01/solucion-apuesta-total/contracts v1.0.0
	github.com/rodsil01/solucion-apuesta-total/messagebroker v1.0.0
)

require github.com/streadway/amqp v1.1.0 // indirect

replace (
	github.com/rodsil01/solucion-apuesta-total/contracts => ../contracts
	github.com/rodsil01/solucion-apuesta-total/messagebroker => ../message-broker
)
