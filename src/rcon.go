package main

import (
	"strconv"

	"github.com/majormjr/rcon"
)

func connectRC(config Config) (*rcon.RemoteConsole, error) {
	rconAddr := config.ServerIP + ":" + strconv.Itoa(config.FactorioRconPort)
	return rcon.Dial(rconAddr, config.FactorioRconPass)
}
