package cmd

import "flag"

func getParams() int {
	port := flag.Int("p", 3000, "Port to serve.")

	flag.Parse()

	return *port
}
