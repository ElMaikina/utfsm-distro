# syntax=docker/dockerfile:1
FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Le pasa al contenedor los archivos para Buildear
COPY go.mod go.sum ./
RUN mkdir ./proto

# Actualiza y descarga las dependencias
RUN apt-get update
RUN apt-get -y install autoconf automake libtool curl make g++ unzip protobuf-compiler

# Entra en la carpeta, genera el servidor gRPC y sale de la carpeta
COPY proto/director.pb.go proto/director_grpc.pb.go ./proto

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
#RUN CGO_ENABLED=0 GOOS=linux go build -o /gemaestrategica

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8000

# Run
CMD ["go", "run", "tierra.go"]