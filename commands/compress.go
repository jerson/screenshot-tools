package commands

import (
	"fmt"
	"os"
	"sort"

	log "github.com/sirupsen/logrus"

	"gopkg.in/AlecAivazis/survey.v1"

	"screenshot-tools/modules/utils"
)

// CompressOptions ...
type CompressOptions struct {
	Input  string
	Output string
}

// Compress ...
func Compress(options CompressOptions) error {

	scenarios, err := ReportDir(options.Input)
	if err != nil {
		return err
	}

	caseGroup := map[string]Case{}
	for _, scenario := range scenarios {
		for _, caseItem := range scenario.Cases {
			key := fmt.Sprintf("%s_%s (%d steps)", caseItem.Name, caseItem.Device, len(caseItem.Steps))
			caseGroup[key] = caseItem

		}
	}

	var caseOptions []string
	for key := range caseGroup {
		caseOptions = append(caseOptions, key)
	}
	sort.Strings(caseOptions)

	var caseOptionsSelected []string
	prompt := &survey.MultiSelect{
		Message:  "Choose cases",
		Options:  caseOptions,
		PageSize: 10,
		Default:  caseOptions,
	}
	err = survey.AskOne(prompt, &caseOptionsSelected, nil)
	if err != nil {
		return err
	}

	for i, caseOption := range caseOptionsSelected {
		caseItem := caseGroup[caseOption]

		log.Infof("[%d/%d] processing: %s_%s (%d steps)", i, len(caseOptionsSelected), caseItem.Name, caseItem.Device, len(caseItem.Steps))

		var filePaths []string
		for _, step := range caseItem.Steps {
			filePaths = append(filePaths, step.Path)
		}

		_ = os.MkdirAll(options.Output, 0777)
		output := fmt.Sprintf("%s/%s_%s.png", options.Output, caseItem.Device, caseItem.Name)
		err := utils.MergeImages(filePaths, output)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Infof("output: %s", output)

	}

	return nil
}
