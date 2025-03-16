package cli

import (
	"slices"

	"github.com/wraith29/apollo/pkg/config"
)

type commandFunc func([]string, *config.Config) error

type Command struct {
	name    string
	exec    commandFunc
	aliases []string
}

func NewCommand(name string, exec commandFunc, aliases ...string) Command {
	return Command{
		name,
		exec,
		aliases,
	}
}

func (c *Command) matches(cmd string) bool {
	if cmd == c.name {
		return true
	}

	return slices.Contains(c.aliases, cmd)
}
