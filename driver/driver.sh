#!/bin/bash

while `true`; do
  seed=$RANDOM$RANDOM
  echo $seed
  generate -dir ./ -seed $seed -out a.go

  GOROOT=$HOME/go1.4 $HOME/go1.4/bin/go run a.go &> good
  if [ $? -ne 0 ]; then
    echo =================================
    cat a.go
    echo ---------------------------------
    cat good
    echo =================================
    continue
  fi

  GOROOT=$HOME/go $HOME/go/bin/go run a.go &> bad
  diff -q -w good bad
  if [ $? -ne 0 ]; then
    exit 1
  fi
done
