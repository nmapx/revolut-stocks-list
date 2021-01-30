#!/usr/bin/env sh

set -e

if [ "$1" = "run" ]
  then exec go run .
elif [ "$1" = "loop" ]
  then exec sh -c "while true; do echo ping; sleep 1; done"
fi

exec $@
