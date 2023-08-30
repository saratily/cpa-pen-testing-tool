package main

import (
	"cpa-pen-testing-tool/internal/cli"
	"cpa-pen-testing-tool/internal/conf"
	"cpa-pen-testing-tool/internal/server"
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}
