package main

import (
	"flag"
	"fmt"
	"github.com/arran4/trace-redirects"
	"log"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Printf("Please provide some URLs\n")
		return
	}

	for _, u := range flag.Args() {
		log.Printf("Tracing %s\n", u)
		ls, err := traceredirects.Trace(u)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		for n, l := range ls {
			fmt.Printf("%d: %s\n", n+1, l)
		}
	}
	fmt.Printf("Done\n")
}
