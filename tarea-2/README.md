# Laboratorio 3 - Sistemas Distribuidos

## Grupo 2
- Leandro Espinoza 201973563-2
- Miguel Soto 201973623-K

## Tutoriales usados:

- Instalacion: https://www.youtube.com/watch?v=gbrPMv_GuQY
- Codigo de Ejemplo: https://www.youtube.com/watch?v=BdzYdN_Zd9Q

## Dependencias

```
go get -u google.golang.org/protobuf/reflect/protoreflect
go get -u google.golang.org/protobuf/runtime/protoimpl
go get -u reflect
go get -u sync
```

## Instrucciones

Para ejecutar mediante _Docker_

```
cd tierra
docker build --tag gemaestrategica .
docker image ls
docker run -p 80:8000 gemaestrategica
```

Para ejecutar de manera local, primero ejecutar en un terminal desde la carpeta del proyecto lo siguiente:

```
cd tierra
go run tierra.go
```

En otro terminal separado (partiendo en la carpeta del proyecto) ejecutar en un terminal lo siguiente:

```
cd equipo
go run equipo.go
```