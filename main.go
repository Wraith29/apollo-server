package main

import (
	"fmt"
	"os"

	"github.com/wraith29/apollo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
