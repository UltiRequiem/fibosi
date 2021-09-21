package cmd

import "github.com/UltiRequiem/fibonnaci/internal"

func Exec() {
	port := getParams()

	internal.NewServer(port)
}
