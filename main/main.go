package main

import (
	"github.com/rs/zerolog/log"
	"github.com/skarfie123/dockertest/internal"
)

func main() {
	println(internal.GetMessage("Rahul"))
	log.Info().Msg("Printed Message")
}
