package main

import (
	"fmt"
	"os"

	"github.com/jutkko/mindown/input"
	"github.com/jutkko/mindown/output"
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
		cli.StringFlag{
			Name:  "output-file",
			Value: "output.txt",
			Usage: "output file name",
		},
		cli.BoolFlag{
			Name:  "f",
			Usage: "provide this flag to overwrite the output file if it exists",
		},
	}

	app.Action = func(c *cli.Context) error {
		if err != nil {
			return err
		}

		err = output.WriteMarkdown(c.String("output-file"), c.Bool("f"), graph)
		if err != nil {
			return err
		}

		return nil
	}

	app.Run(os.Args)
}
