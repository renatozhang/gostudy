package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args().Get(0)
		}
		fmt.Println("hello firend! cmd:", cmd)
		return nil
	}
	app.Run(os.Args)
}
