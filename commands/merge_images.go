package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"screenshot_tools/modules/utils"
)

// MergeImagesOptions ...
type MergeImagesOptions struct {
	Input      string
	OutputFile string
}

// MergeImages ...
func MergeImages(options MergeImagesOptions) error {

	files, err := ioutil.ReadDir(options.Input)
	if err != nil {
		return err
	}
	var filePaths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := fmt.Sprintf("%s/%s", options.Input, file.Name())
		extension := filepath.Ext(path)
		if !contains([]string{".jpg", ".png"}, strings.ToLower(extension)) {
			log.Warnf("ignored file: %s", path)
			continue
		}
		filePaths = append(filePaths, path)

	}

	if len(filePaths) < 1 {
		return errors.New("images not found")
	}

	err = utils.MergeImages(filePaths, options.OutputFile)
	if err != nil {
		return err
	}
	log.Infof("output: %s\n", options.OutputFile)

	return nil
}
