package cli

import (
	"flag"
	"fmt"
	"os"

	"cpa-pen-testing-tool/internal/logging"
)

func usage() {
	fmt.Print(`This program runs CPA pen testing tool backend server.

Usage:

cpa-pen-testing-tool [arguments]

Supported arguments:

`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() string {
	flag.Usage = usage
	env := flag.String("env", "dev", `Sets run environment. Possible values are "dev" and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetGinLogToFile()
	}
	return *env
}
