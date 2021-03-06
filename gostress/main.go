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

	m := &fnct{nam: "main", typ: ""}
	f := newFnct(choice(basicTypes...))
	c := m.call(f)

	if *want == "" {
		fmt.Fprintf(w, "// gostress -seed %d\n", *seed)
		fmt.Fprintf(w, "package main\n")
		fmt.Fprintf(w, "import \"fmt\"\n")
		fmt.Fprintf(w, "func b2i(b bool) int {\n")
		fmt.Fprintf(w, "if b {\n")
		fmt.Fprintf(w, "return 1\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "return 0\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "func main() {\n")
		fmt.Fprintf(w, "fmt.Println(%s)\n", c)
		fmt.Fprintf(w, "}\n")
	} else {
		fmt.Fprintf(w, "// gostress -seed %d -want %s\n", *seed, *want)
		fmt.Fprintf(w, "package main\n")
		fmt.Fprintf(w, "import \"fmt\"\n")
		fmt.Fprintf(w, "import \"os\"\n")
		fmt.Fprintf(w, "func b2i(b bool) int {\n")
		fmt.Fprintf(w, "if b {\n")
		fmt.Fprintf(w, "return 1\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "return 0\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "func main() {\n")
		fmt.Fprintf(w, "if got := %s; got != %s {\n", c, *want)
		fmt.Fprintf(w, "fmt.Printf(\"%s = %%v, wanted %s\\n\", got)\n", c, *want)
		fmt.Fprintf(w, "os.Exit(1)\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "}\n")
	}

	f.dump(w)
}
