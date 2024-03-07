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
3. Servicio de Reservas: Servicio encargado de realizar reservas de vuelos
4. Servicio de Equipajes: Servicio encargado de administrar la información de equipajes de las reservas realizadas
