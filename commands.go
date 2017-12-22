package main
// https://github.com/home-assistant/hassio/blob/dev/API.md
import (
	"fmt"
	"os"

	"github.com/home-assistant/hassio-cli/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "homeassistant",
		Usage:  "",
		Action: command.CmdHomeassistant,
		Flags:  []cli.Flag{
			cli.BoolFlag{
				Name: "rawjson",
				Usage: "Returns the output in JSON format",
			},
			cli.StringFlag{
				Name: "payload",
				Usage: "holds data for POST in format key=val,key2=val2",
			},
		},
	},
	{
		Name:   "supervisor",
		Usage:  "",
		Action: command.CmdSupervisor,
		Flags:  []cli.Flag{
			cli.BoolFlag{
				Name: "rawjson",
				Usage: "Returns the output in JSON format",
			},
			cli.StringFlag{
				Name: "payload",
				Usage: "holds data for POST in format key=val,key2=val2",
			},
		},
	},
	{
		Name:   "host",
		Usage:  "",
		Action: command.CmdHost,
		Flags:  []cli.Flag{
			cli.BoolFlag{
				Name: "rawjson",
				Usage: "Returns the output in JSON format",
			},
			cli.StringFlag{
				Name: "payload",
				Usage: "holds data for POST in format key=val,key2=val2",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
