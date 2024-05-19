# Laboratorio 3 - Sistemas Distribuidos

## Grupo 2
- Leandro Espinoza 201973563-2
- Miguel Soto 201973623-K

## Dependencias

```
go get -u "github.com/rabbitmq/amqp091-go"
```

## Instrucciones

Se asume de antemano que las máquinas están corriendo y que el servicio de _Docker_ está activo

Para ejecutar el director, hacer _SSH_ a la siguiente IP
´´´
ssh dist@dist006.inf.santiago.usm.cl
cd utfsm-distro/tarea-3/director
sudo docker build --tag director .
sudo docker run -p 50051:50051 director
´´´