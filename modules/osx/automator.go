package osx

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	"github.com/gobuffalo/packr/v2"

	"screenshot_tools/modules/config"
	"screenshot_tools/modules/utils"
)

// OpenApp ...
func OpenApp(name string) error {

	err := exec.Command("open", "-a", name).Run()
	if err != nil {
		return err
	}
	return nil
}

// GetAutomatorDir ...
func GetAutomatorDir() (string, error) {
	output := ""
	usr, err := user.Current()
	if err != nil {
		return output, err
	}
	output = fmt.Sprintf("%s/%s", usr.HomeDir, config.Vars.Dir.Automator)
	return output, nil
}

// GetAutomatorFile ...
func GetAutomatorFile(name string) (string, error) {
	output := ""
	dir, err := GetAutomatorDir()
	if err != nil {
		return output, err
	}
	output = fmt.Sprintf("%s/assets/automator/%s", dir, name)
	return output, nil
}

// LoadAssets ...
func LoadAssets(box *packr.Box) error {
	bytes, err := box.Find("automator.zip")
	if err != nil {
		return err
	}

	dir, err := GetAutomatorDir()
	if err != nil {
		return err
	}
	_ = os.Mkdir(dir, 0777)
	_, err = utils.Unzip(bytes, dir)
	return err
}
