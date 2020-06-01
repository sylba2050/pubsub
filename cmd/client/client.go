package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	ps "github.com/sylba2050/pubsub"
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
		text, _ := reader.ReadBytes('\n')

		payload, _ := ps.NewPayload(ps.MessageBody)
		payload.SetLength(uint16(len(text)))
		payload.SetValue(text)
		dataLength, _ := payload.GetLength()

		header, _ := ps.NewHeader(ps.Publish)
		// NewMessageを改造して自動でheaderにlengthが設定されるように
		header.SetLength(dataLength + ps.PayloadHeaderSize)
		header.SetReceiverTimestamp(1)
		header.SetSenderTimestamp(1)

		message := ps.NewMessage(header, payload)
		fmt.Printf("%+v\n", payload)
		fmt.Printf("%+v\n", header)
		b, _ := message.ToBytes()
		fmt.Fprintf(c, string(b))

		receive, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + receive)
		if strings.TrimSpace(string(text)) == "STOP" {
			log.Info().Msg("TCP client exiting...")
			return
		}
	}
}
