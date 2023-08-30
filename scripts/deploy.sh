#! /bin/bash

# default ENV is dev
env=dev

while test $# -gt 0; do
  case "$1" in
    -env)
      shift
      if test $# -gt 0; then
        env=$1
      fi
      # shift
      ;;
    *)
    break
    ;;
  esac
done

cd ../../cpa-pen-testing-tool
source .env
go build -o cmd/cpa-pen-testing-tool/cpa-pen-testing-tool cmd/cpa-pen-testing-tool/main.go
cmd/cpa-pen-testing-tool/cpa-pen-testing-tool -env $env &
