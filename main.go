package main

import (
	"fmt"
	"os"

	"github.com/jutkko/mindown/input"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mindown"
	app.Usage = "convert mind to files"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input-file",
			Value: "input.txt",
			Usage: "input file name",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Input file name: %s\n", c.String("input-file"))
		graph, err := input.ParseOpml(c.String("input-file"))
		if err != nil {
			panic(err.Error())
		}

		err = graph.Export()

		if err != nil {
			panic(err.Error())
		}

		return nil
	}

	app.Run(os.Args)
}
