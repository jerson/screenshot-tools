package actions

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"

	"screenshot_tools/commands"
	"screenshot_tools/modules/config"
	"screenshot_tools/modules/prompt"
	"screenshot_tools/modules/utils"
)

// MergeImages ...
func MergeImages(c *cli.Context) error {
	input := c.String("input")
	output := c.String("output")

	input = prompt.Dir("Input Dir", input, config.Vars.Dir.Input, true)

	suggestion, err := utils.GetEvidenceSuggestion(input)
	if err != nil {
		log.Warn(err)
	}
	var names []string
	if suggestion.Model != "" {
		names = append(names, suggestion.Model)
	}
	if suggestion.Name != "" {
		names = append(names, suggestion.Name)
	}
	if len(names) < 1 {
		names = append(names, "merged")
	}

	output = prompt.Field("Output File", output, "", fmt.Sprintf("%s.png", strings.Join(names, "_")))

	options := commands.MergeImagesOptions{
		Input:      input,
		OutputFile: output,
	}
	return commands.MergeImages(options)
}
