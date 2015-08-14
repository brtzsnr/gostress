#!/bin/bash

while `true`; do
  # generate a random program and record the value.
  seed=$RANDOM$RANDOM
  generate -seed $seed -out tmp.a$$.go

  GOROOT=$HOME/go1.4 $HOME/go1.4/bin/go build tmp.a$$.go 2> /dev/null
  if [ $? -ne 0 ]; then
    #echo =================================
    #cat a$$.go
    #echo =================================
    continue
  fi

  # generate the same program but now it tests the result.
  ./tmp.a$$ > tmp.good$$
  generate -seed $seed -out tmp.a$$.go -want `cat tmp.good$$`

  GOROOT=$HOME/go $HOME/go/bin/go build tmp.a$$.go 2> tmp.log$$
  if [ $? -ne 0 ]; then
    echo broken $seed
    tail -n1 tmp.log$$
    continue
  fi

  ./tmp.a$$
  if [ $? -ne 0 ]; then
    echo failed $seed
    cp tmp.a$$.go failed-$seed.go
    continue
  fi
done
