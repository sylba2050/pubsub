package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	ps "github.com/sylba2050/pubsub"
)

func handle(c net.Conn) {
	defer c.Close()

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Error().Err(err).Send()
			return
		}

		if strings.TrimSpace(netData) == "STOP" {
			log.Info().Msg("close connection")
			return
		}

		fmt.Print("-> ", netData)
		myTime := time.Now().Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
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

		go handle(c)
	}
}
