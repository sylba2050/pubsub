package pubsub

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// Config is instance of configuration
var Config config

type config struct {
	Port    int    `yaml:"port"`
	LogType string `yaml:"log_type"`
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func init() {
	parse()
}

func parse() {
	c := config{}
	var err error

	data, err := ioutil.ReadFile("./config.yaml")
	log.Info().Msg(string(data))
	if err != nil {
		log.Fatal().Msg("Can't read configuration file")
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatal().Msg("Can't parse configuration file")
	}
	Config = c
}
