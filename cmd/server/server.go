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

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("%d", ps.Config.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	defer c.Close()

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
