package main

import (
	"flag"
	"log"
	"trace-redirects"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Printf("Please provide some URLs")
		return
	}

	for _, u := range flag.Args() {
		log.Printf("Tracing %s", u)
		ls, err := traceredirects.Trace(u)
		if err != nil {
			log.Printf("Error: %s", err)
		}
		for n, l := range ls {
			log.Printf("%d: %s", n+1, l)
		}
	}
	log.Printf("Done")
}
