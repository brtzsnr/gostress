#!/bin/bash

while `true`; do
  seed=$RANDOM$RANDOM
  generate -dir ./ -seed $seed -out a$$.go

  GOROOT=$HOME/go1.4 $HOME/go1.4/bin/go build a$$.go 2> /dev/null
  if [ $? -ne 0 ]; then
    #echo =================================
    #cat a.go
    #echo =================================
    continue
  fi

  ./a$$ &> good$$

  GOROOT=$HOME/go $HOME/go/bin/go build a$$.go 2> log$$
  if [ $? -ne 0 ]; then
    echo broken $seed
    tail -n1 log$$

    printf "// " > broken-$seed.go
    tail -n1 log$$ >> broken-$seed.go
    cat a$$.go >> broken-$seed.go
    continue
  fi
  ./a$$ &> bad$$

  diff -q -w good$$ bad$$
  if [ $? -ne 0 ]; then
    echo failed $seed
    print good `cat good$$`
    print bad `cat bad$$`

    cp a$$.go failed-$seed.go
    continue
  fi
done
