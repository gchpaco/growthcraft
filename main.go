package main

import "flag"
import "log"
import "fmt"
import "os"
import "github.com/gchpaco/growthcraft/parser"

var rootdir = flag.String("d", "", "directory to search for files")

func main() {
	flag.Parse()

	if *rootdir == "" {
		log.Fatal("Cannot proceed without a root directory")
	}

	fmt.Println("digraph Growthcraft {")
	fmt.Println("\tnull [style=invis,rank=source]")
	fmt.Println("\trankdir=LR")

	{
		f, err := os.Open(*rootdir + "/config/growthcraft/cellar/brewing.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		brews, err := parser.DecodeBrewing(f)
		if err != nil {
			log.Fatal(err)
		}

		for _, brew := range brews {
			brew.Visit(os.Stdout)
		}
	}

	{
		f, err := os.Open(*rootdir + "/config/growthcraft/cellar/fermenting.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		ferments, err := parser.DecodeFerment(f)
		if err != nil {
			log.Fatal(err)
		}

		for _, ferment := range ferments {
			ferment.Visit(os.Stdout)
		}
	}

	{
		f, err := os.Open(*rootdir + "/config/growthcraft/cellar/pressing.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		presses, err := parser.DecodePressing(f)
		if err != nil {
			log.Fatal(err)
		}

		for _, press := range presses {
			press.Visit(os.Stdout)
		}
	}

	for key, value := range parser.Fluids {
		fmt.Printf("\t\"%s\"[label=\"%s\"];\n", key, value)
	}

	// Dirty hacks
	fmt.Println(`
	"grc.applecider2"->"grc.applecider3"[arrowhead=normal, label="Dust     ", style=solid, labeljust="r", dir=both, constraint=false];
	"grc.grapewine2"->"grc.grapewine3"[arrowhead=normal, label="Dust     ", style=solid, labeljust="r", dir=both, constraint=false];
	"grc.hopale2"->"grc.hopale3"[arrowhead=normal, label="Dust     ", style=solid, labeljust="r", dir=both, constraint=false];
	"grc.lager2"->"grc.lager3"[arrowhead=normal, label="Dust     ", style=solid, labeljust="r", dir=both, constraint=false];
	"grc.ricesake2"->"grc.ricesake3"[arrowhead=normal, label="Dust     ", style=solid, labeljust="r", dir=both, constraint=false];
	"grc.honeymead2"->"grc.honeymead3"[arrowhead=normal, label="Dust     ", style=solid, labeljust="r", dir=both, constraint=false];
`)

	fmt.Println("}")
}
