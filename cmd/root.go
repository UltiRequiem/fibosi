package cmd

import "github.com/UltiRequiem/fibonnaci/internal"

func Exec() {
	port, banner := getParams()

	internal.NewServer(port, banner)
}
