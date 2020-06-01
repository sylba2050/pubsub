package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/rs/zerolog/log"

	ps "github.com/sylba2050/pubsub"
)

func handleClient(c net.Conn) {
	defer c.Close()

	isAuthorized, err := ps.IsAuthorized()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	if !isAuthorized {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		listen(c)
	}()

	go func() {
		defer wg.Done()
		sendDataToClient(c)
	}()
	wg.Wait()
}

func listen(c net.Conn) {
	for {
		header, err := ps.ReadHeader(c)
		if err != nil {
			return
		}
		// FIXME headerのサイズを元に繰り返しpayload読み取り
		payload, err := ps.ReadPayload(c)
		if err != nil {
			return
		}
		handleReceive(header, payload)
	}
}

func handleReceive(header ps.Header, payload ps.Payload) error {
	// TODO Implement
	// TODO payloadを可変長引数に
	fmt.Printf("Header: %+v\n", header)
	fmt.Printf("Payload: %+v\n", payload)
	return nil
}

func sendDataToClient(c net.Conn) {
	// TODO Implement
	for {
		fmt.Fprintf(c, "Hello, World\n")
	}
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", ps.Config.Port))
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		log.Info().Msg("accept")
		if err != nil {
			log.Error().Err(err).Send()
			return
		}

		go handleClient(c)
	}
}
