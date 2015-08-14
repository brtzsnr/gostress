package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"math/rand"
)

var (
	seed = flag.Int64("seed", 1, "random seed")
	out = flag.String("out", "-", "output file")
)

func main() {
	flag.Parse()
	rand.Seed(*seed)

	var w *os.File
	if *out == "-" {
		w = os.Stdout
	} else {
		var err error
		w, err = os.Create(*out)
		if err != nil {
			log.Fatal("Create: ", err)
		}
		defer w.Close()
	}

	f := newFun(choice(basicTypes...))
	fmt.Fprintf(w, "package main\nimport \"fmt\"\n")
	fmt.Fprintf(w, "func main() {\nfmt.Println(%s())\n}\n", f.nam)
	f.dump(w)
}
