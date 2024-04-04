package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	// Se define la direccion a la cual se va a
	// conectar para enviar informacion al servidor
	// central
	IP := "172.17.0.2:"
	PORT := "8081"

	captains := []string{"CA", "CB", "CC", "CD", "CE", "CF"}
	planets := []string{"PA", "PB", "PC", "PD", "PE", "PF"}

	udpAddr := IP + PORT
	addr, _ := net.ResolveUDPAddr("udp", udpAddr)

	for true {
		delay := rand.Int31n(1000)
		time.Sleep(time.Duration(delay)*time.Millisecond + 100)
		conn, err := net.DialUDP("udp", nil, addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		captain := captains[rand.Intn(6)]
		planet := planets[rand.Intn(6)]
		bounty := rand.Intn(100)
		msg := captain + "," + planet + "," + strconv.Itoa(bounty)

		fmt.Println("Capitán", captain, "encontró botín en", planet, "enviando solicitud de asignacion")

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error: ", err)
		}

	}
}
