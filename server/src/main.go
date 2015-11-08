package main

import (
	"ts2"
	"fmt"
	"encoding/json"
	"flag"
)

func main() {
	simJson := flag.String("load", "", "A JSON string with the simulation definition to load")
	flag.Parse()
	var sim ts2.Simulation
	json.Unmarshal([]byte(*simJson), &sim)
	fmt.Println("Press Ctrl-C to end")
	for {
	}
}