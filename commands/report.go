package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// ReportOptions ...
type ReportOptions struct {
	Input string
}

// Scenario ...
type Scenario struct {
	DirName string
	Name    string
	Path    string
	Cases   []Case
}

// Case ...
type Case struct {
	Name   string
	Device string
	Steps  []Step
}

// Step ...
type Step struct {
	FileName string
	Name     string
	Path     string
}

// Report ...
func Report(options ReportOptions) ([]Scenario, error) {
	return ReportDir(options.Input)
}

// ReportDir ...
func ReportDir(baseDir string) ([]Scenario, error) {
	var scenarios []Scenario

	dirs, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		path := fmt.Sprintf("%s/%s", baseDir, dir.Name())
		cases, err := readCases(path)
		if err != nil {
			log.Error(err)
			continue
		}

		scenarios = append(scenarios, Scenario{
			Name:    dir.Name(),
			DirName: dir.Name(),
			Path:    path,
			Cases:   cases,
		})
	}

	if len(scenarios) < 1 {
		log.Info("fallback now using input as scenario")
		cases, err := readCases(baseDir)
		if err != nil {
			log.Error(err)
		} else {
			scenarios = append(scenarios, Scenario{
				Name:    "",
				DirName: "",
				Path:    baseDir,
				Cases:   cases,
			})
		}
	}

	return scenarios, nil
}

func readCases(dir string) ([]Case, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	caseGroup := map[string]*Case{}
	for _, file := range files {
		path := fmt.Sprintf("%s/%s", dir, file.Name())

		if file.IsDir() {
			cases, err := readCases(path)
			if err != nil {
				log.Error(err)
				continue
			}

			for _, caseItem := range cases {

				key := fmt.Sprintf("%s_%s", caseItem.Device, caseItem.Name)
				if caseGroup[key] == nil {
					caseGroup[key] = &Case{
						Name:   caseItem.Name,
						Device: caseItem.Device,
						Steps:  caseItem.Steps,
					}
				} else {
					caseGroup[key].Steps = append(caseGroup[key].Steps, caseItem.Steps...)
				}

			}
			continue
		}

		extension := filepath.Ext(path)
		if !contains([]string{".jpg", ".png"}, strings.ToLower(extension)) {
			log.Warnf("ignored file: %s", path)
			continue
		}

		step := ""
		fileParts := strings.Split(file.Name(), "_")
		if len(fileParts) < 2 {

			dirParts := strings.Split(dir, "/")
			currentDir := dirParts[len(dirParts)-1]
			fileParts = strings.Split(currentDir, "_")
			if len(fileParts) < 2 {
				log.Errorf("invariant name file: %s", path)
				continue
			}
			step = file.Name()
			step = strings.Replace(step, extension, "", 1)
		}

		device := strings.TrimSpace(fileParts[0])
		device = strings.Replace(device, extension, "", 1)
		name := strings.TrimSpace(fileParts[1])
		name = strings.Replace(name, extension, "", 1)
		key := fmt.Sprintf("%s_%s", device, name)
		if caseGroup[key] == nil {
			caseGroup[key] = &Case{
				Name:   name,
				Device: device,
				Steps:  []Step{},
			}
		}

		if len(fileParts) > 2 {
			step = strings.TrimSpace(fileParts[2])
			step = strings.Replace(step, extension, "", 1)
		}

		caseGroup[key].Steps = append(caseGroup[key].Steps, Step{
			FileName: file.Name(),
			Name:     step,
			Path:     path,
		})
	}

	var cases []Case
	for _, caseItem := range caseGroup {
		cases = append(cases, *caseItem)
	}
	return cases, err
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
