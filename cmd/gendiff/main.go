package main

import (
	"context"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

// RunCLI runs the gendiff command with the provided arguments and output writer.
func RunCLI(ctx context.Context, args []string, writer io.Writer) error {
	cmd := &cli.Command{
		Name:                   "gendiff",
		Usage:                  "Compares two configuration files and shows a difference.",
		UsageText:              "gendiff [global options]",
		Writer:                 writer,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
	return cmd.Run(ctx, args)
}

func main() {
	RunCLI(context.Background(), os.Args, os.Stdout)
}
