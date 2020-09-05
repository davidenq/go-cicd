package config

import (
	"log"
	"os"
)

//PORT .
var PORT string

//Load .
func Load() {

	PORT = os.Getenv("PORT")

	if PORT == "" {
		log.Fatalln("PORT is required")
	}
}
