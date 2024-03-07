# Solución

## Consideraciones

- La solución utiliza Docker. La guía mostrada a continuación muestra como iniciar los servicios con docker compose
- Recomiendo usar Postman para realizar las pruebas

## Setup

1. `git clone `
2. `cd solucion-prueba-tecnica`
3. `docker-compose up -d`

## Servicios

La solución consiste en los siguientes servicios:

1. MySQL: Base de datos utilizada para la persistencia de datos
2. RabbitMQ: Servicio de mensajería para la comunicación entre los microservicios
3. Servicio de Reservas  (localhost:8080): Servicio encargado de realizar reservas de vuelos
4. Servicio de Equipajes (localhost:8081): Servicio encargado de administrar la información de equipajes de las reservas realizadas

## Enpoints

### Clientes

GET http://localhost:8080/clients


GET http://localhost:8080/clients/{client_id}

### Rutas

GET http://localhost:8080/routes


GET http://localhost:8080/routes/{route_id}

### Reservas

GET http://localhost:8080/reservations


GET http://localhost:8080/reservations/{reservation_id}




POST http://localhost:8080/reservations

### Equipaje

GET http://localhost:8081/baggage/reservations/{reservation_id}

---

*Se creó un script que initializa las bases de datos utilizadas y agrega valores para facilitar las pruebas*

## Ejemplo de petición de reserva

POST http://localhost:8080/reservations


body:

```json
{
    "routeId": "96de053d-fcfb-4409-8b96-0f7b136e62e0",
    "clientId": "4bc90533-ec62-45b3-9e43-6f1f79b239f0",
    "reservationDate": "2020-07-30T18:00:00.000Z",
    "seats": 2,
    "state": 23,
    "baggage": [
        {
            "description": "equipaje 1",
            "weight": "10.5"
        },
        {
            "description": "equipaje 2",
            "weight": "11.5"
        },
        {
            "description": "equipaje 3",
            "weight": "12.5"
        }
    ]
}
```


