#!/usr/bin/env sh

set -e

if [ "$1" = "run" ]
  then exec make -f Makefile.native run
elif [ "$1" = "loop" ]
  then exec sh -c "while true; do echo ping; sleep 1; done"
fi

exec $@
