package main

import (
	"fmt"
	"os"

	"github.com/wraith29/apollo/pkg/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 1 argument but found 0")
		os.Exit(1)
	}

	cli, err := cli.NewCli(
		os.Args[1:],
		cli.NewCommand("start", start),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cli.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
