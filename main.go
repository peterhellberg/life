package main

import (
	"flag"

	"github.com/peterhellberg/life/life"
)

var variant = flag.String("variant", "life", "Variations on the rules [life, daynight, highlife, seed]")

func main() {
	flag.Parse()

	life.Run(variant)
}
