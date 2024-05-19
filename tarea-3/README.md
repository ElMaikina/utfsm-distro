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

Para correr el Data Node ejecutar los siguientes comandos
´´´
ssh dist@dist005.inf.santiago.usm.cl
cd utfsm-distro/tarea-3/data_node
sudo docker build --tag data_node .
sudo docker run -p 50053:50053 data_node
´´´

Para correr el Director ejecutar los siguientes comandos
´´´
ssh dist@dist006.inf.santiago.usm.cl
cd utfsm-distro/tarea-3/director
sudo docker build --tag director .
sudo docker run -p 50051:50051 director
´´´

Para correr el Name Node ejecutar los siguientes comandos
´´´
ssh dist@dist008.inf.santiago.usm.cl
cd utfsm-distro/tarea-3/name_node
sudo docker build --tag name_node .
sudo docker run -p 50052:50052 name_node
´´´