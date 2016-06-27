package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
	"path"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "A distributed lock management CLI and library for Docker platforms."
	app.Version = VERSION
	app.Authors = []cli.Author{{Name: "Jeff Nickoloff", Email: "jeff@allingeek.com"}}
	app.Flags = flags
	app.Commands = commands
	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level)
		if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func commandNotYetImplemented(c *cli.Context) error {
	fmt.Printf("Invocation: %s [xxx] %s\n", c.Command.Name, c.Args())
	fmt.Println("This command has yet to be implemented.")
	return nil
}

var (
	flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (options: debug, info, panic)"),
		},
		cli.StringSliceFlag{
			Name:   "host, H",
			EnvVar: "DOCKER_HOST",
			Value:  &cli.StringSlice{"tcp://:2376"},
			Usage:  "Docker API endpoint.",
		},
	}
	commands = []cli.Command{
		{
			Name:  "acquire",
			Usage: "Acquire and hold a named lock.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image, i",
					Usage: "The lock image to use.",
					Value: "alpine",
				},
			},
			Before: func(c *cli.Context) error {
				args := c.Args()
				if len(args) != 1 {
					return errors.New("Lock name is required.")
				}
				return nil
			},
			Action: commandNotYetImplemented,
		},
		{
			Name:  "interrupt",
			Usage: "Remove failure policy",
			Before: func(c *cli.Context) error {
				args := c.Args()
				if len(args) != 1 {
					return errors.New("Lock name is required.")
				}
				return nil
			},
			Action: commandNotYetImplemented,
		},
		{
			Name:   "ls",
			Usage:  "List locks",
			Action: commandNotYetImplemented,
		},
		{
			Name:   "info",
			Usage:  "Retrive information about the underlying Docker system.",
			Action: commandNotYetImplemented,
		},
		{
			Name:      "version",
			ShortName: "v",
			Usage:     "Show version",
			Action: func(c *cli.Context) error {
				PrintVersion()
				return nil
			},
		},
	}
)
