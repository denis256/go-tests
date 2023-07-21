package main

import (
	"log"

	"noerrorcheck/config"
)

func main() {
	c, err := config.CreateEvalConfig("error")
	log.Printf("ctx: %s, err: %v\n", c, err)

	c, err = config.CreateEvalConfig("data")
	log.Printf("ctx: %s, err: %v\n", c, err)
}
