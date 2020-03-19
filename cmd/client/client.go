package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatal().Msg("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		log.Fatal().Err(err).Send()
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			log.Info().Msg("TCP client exiting...")
			return
		}
	}
}
