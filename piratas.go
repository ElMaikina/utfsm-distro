package main

import (
	"fmt"
	"net"
    "math/rand"
	"time"
)

func main() {
	// Se define la direccion a la cual se va a
	// conectar para enviar informacion al servidor
	// central
	IP := "localhost:"
	PORT := "8081"

	direccion_servidor := IP + PORT
	direccion_udp, error_resolver := net.ResolveUDPAddr("udp", direccion_servidor)

	if error_resolver != nil {
		fmt.Println("Error: ", error_resolver)
	}

	// Conexión con el servidor
	conexion, error_al_escuchar := net.ListenUDP("udp4", direccion_udp)
	if error_al_escuchar != nil {
		fmt.Println("Error:", error_al_escuchar)
	}
	conexion.Close()

	fmt.Println("Conexión establecida exitosamente!")
	rand.Seed(time.Now().Unix())

	for true {
		yourTime := rand.Int31n(1000)
		time.Sleep(time.Duration(yourTime) * time.Millisecond + 100)
		conexion_de_envio, err := net.DialUDP("udp", nil, direccion_udp)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		_, err = conexion_de_envio.Write([]byte("1,1,1"))
		fmt.Println("Enviando mensaje...")

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
