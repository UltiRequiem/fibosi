// Provides an easy way to initialize the server
package cmd

import (
	"log"

	. "github.com/UltiRequiem/fibonacci/internal"
)

// Initialize the process
func Exec() {
	port, banner := getParams()

	err := NewServer(port, banner)

	if err != nil {
		log.Fatal(err)
	}
}
