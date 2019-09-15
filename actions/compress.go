package actions

import (
	"github.com/urfave/cli"

	"screenshot-tools/commands"
	"screenshot-tools/modules/config"
	"screenshot-tools/modules/prompt"
)

// Compress ...
func Compress(c *cli.Context) error {
	input := c.String("input")
	output := c.String("output")

	input = prompt.Dir("Input Dir", input, config.Vars.Dir.Input, true)
	output = prompt.Dir("Output Dir", output, config.Vars.Dir.Output, true)

	options := commands.CompressOptions{
		Input:  input,
		Output: output,
	}
	err := commands.Compress(options)
	return err
}
