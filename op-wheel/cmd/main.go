package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/ethereum/go-ethereum/log"

	oplog "github.com/ethereum-optimism/pepe/op-service/log"
	wheel "github.com/ethereum-optimism/pepe/op-wheel"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	app := cli.NewApp()
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "op-wheel"
	app.Usage = "Pepe Wheel is a CLI tool for the execution engine"
	app.Description = "Pepe Wheel is a CLI tool to direct the engine one way or the other with DB cheats and Engine API routines."
	app.Flags = []cli.Flag{wheel.GlobalGethLogLvlFlag}
	app.Before = func(c *cli.Context) error {
		log.Root().SetHandler(
			log.LvlFilterHandler(
				oplog.Level(c.GlobalString(wheel.GlobalGethLogLvlFlag.Name)),
				log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
			),
		)
		return nil
	}
	app.Action = cli.ActionFunc(func(c *cli.Context) error {
		return errors.New("see 'cheat' and 'engine' subcommands and --help")
	})
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Commands = []cli.Command{
		wheel.CheatCmd,
		wheel.EngineCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
