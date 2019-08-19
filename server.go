package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 3000})
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	buff := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Printf("Unable to read from UDP: %s\n", err)
		}
		fmt.Printf("Recived %s, from: %s\n", buff[0:n], addr)
		go sendResponse(conn, addr, buff[0:n])
	}

}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, message []byte) {
	fmt.Printf("Sending message:%s to %s\n", message, addr)
	_, err := conn.WriteToUDP(message, addr)
	if err != nil {
		fmt.Printf("Error while sendind response: %s", err)
	}
}
