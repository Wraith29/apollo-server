package cli

import (
	"errors"

	"github.com/wraith29/apollo/pkg/config"
)

var (
	MissingCommandErr  = errors.New("no command specified")
	InvalidArgumentErr = errors.New("invalid argument")
)

type Cli struct {
	args     []string
	commands []Command
	cfg      *config.Config
}

func NewCli(args []string, commands ...Command) (Cli, error) {
	cfg, err := config.Load()

	return Cli{
		args,
		commands,
		cfg,
	}, err
}

func (c *Cli) Run() error {
	if len(c.args) < 1 {
		return MissingCommandErr
	}

	cmd := c.args[0]

	for _, command := range c.commands {
		if !command.matches(cmd) {
			continue
		}

		return command.exec(c.args[1:], c.cfg)
	}

	return nil
}
