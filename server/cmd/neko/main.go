package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"goldenlumia/neko"
	"goldenlumia/neko/cmd"
	"goldenlumia/neko/internal/utils"
)

func main() {
	fmt.Print(utils.Colorf(neko.Header, "server", neko.Service.Version))
	if err := cmd.Execute(); err != nil {
		log.Panic().Err(err).Msg("failed to execute command")
	}
}
