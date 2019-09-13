package actions

import (
	"encoding/json"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"

	"screenshot_tools/commands"
	"screenshot_tools/modules/config"
	"screenshot_tools/modules/prompt"
)

// Report ...
func Report(c *cli.Context) error {

	input := c.String("input")
	input = prompt.Dir("Input Dir", input, config.Vars.Dir.Input, true)

	options := commands.ReportOptions{
		Input: input,
	}

	data, err := commands.Report(options)
	if data != nil {
		printJSON(data)
	}
	return err
}

func printJSON(data interface{}) {

	output, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	log.Debug(string(output))
}
