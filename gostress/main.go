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
	want = flag.String("want", "", "wanted result")
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

	if *want == "" {
		fmt.Fprintf(w, "// generate -seed %d\n", *seed)
		fmt.Fprintf(w, "package main\n")
		fmt.Fprintf(w, "import \"fmt\"\n")
		fmt.Fprintf(w, "func main() {\n")
		fmt.Fprintf(w, "fmt.Println(%s())\n", f.nam)
		fmt.Fprintf(w, "}\n")
	} else {
		fmt.Fprintf(w, "// generate -seed %d\n", *seed)
		fmt.Fprintf(w, "package main\n")
		fmt.Fprintf(w, "import \"fmt\"\n")
		fmt.Fprintf(w, "import \"os\"\n")
		fmt.Fprintf(w, "func main() {\n")
		fmt.Fprintf(w, "if got := %s(); got != %s {\n", f.nam, *want)
		fmt.Fprintf(w, "fmt.Printf(\"%s() = %%v, wanted %s\\n\", got)\n", f.nam, *want)
		fmt.Fprintf(w, "os.Exit(1)\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "}\n")
	}

	f.dump(w)
}
