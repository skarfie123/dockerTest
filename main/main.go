package main

import (
	"github.com/rs/zerolog/log"
	"github.com/skarfie123/dockertest/msg"
)

func main() {
	println(msg.GetMessage("Rahul"))
	log.Info().Msg("Printed Message")
}
