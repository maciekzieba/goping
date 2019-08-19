package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	service := "localhost:3000"
	remoteAddr, err := net.ResolveUDPAddr("udp", service)
	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	packageID := 0
	buff := make([]byte, 1024)
	for {
		time.Sleep(1000 * time.Millisecond)
		packageID++
		send := time.Now()
		_, err = conn.Write([]byte(fmt.Sprintf("%d", packageID)))
		if err != nil {
			fmt.Printf("Unable to send udp package %d: %s\n", packageID, err)
			continue
		}
		fmt.Printf("Send %d ", packageID)
		n, _, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Printf("unable to read from UDP: %s\n", err)
			continue
		}
		if string(buff[0:n]) == fmt.Sprintf("%d", packageID) {
			receive := time.Now()
			dur := receive.Sub(send)
			fmt.Printf("ping took: %2.3f ms, %s\n", float64(dur)/float64(time.Millisecond), dur.String())
		} else {
			fmt.Printf("error occured, wrong package ID, should be: %d got: %s\n", packageID, buff[0:n])
		}
	}
}
