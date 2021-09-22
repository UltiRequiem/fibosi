package cmd

import "flag"

// Get server configuration
func getParams() (int, bool) {
	port := flag.Int("p", 3000, "Port to serve.")
	banner := flag.Bool("b", false, "Show or Hide ECHO Banner.")

	flag.Parse()

	return *port, *banner
}
