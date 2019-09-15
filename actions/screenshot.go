package actions

import (
	"errors"
	"fmt"
	"time"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"

	"screenshot-tools/commands"
	"screenshot-tools/modules/config"
	"screenshot-tools/modules/prompt"
	"screenshot-tools/modules/utils"
)

// ScreenShot ...
func ScreenShot(c *cli.Context) error {

	name := c.Args().Get(0)
	adb := c.String("adb")
	automator := c.String("automator")
	platform := c.String("platform")

	if !(platform == "android" || platform == "ios" || platform == "ios-simulator" || platform == "desktop") {
		return fmt.Errorf("not implemented for: %s", platform)
	}

	suggestion, err := utils.GetEvidenceSuggestion(".")
	if err != nil {
		log.Warn(err)
	}

	name = prompt.Field("Name", name, "name", suggestion.Model)
	adb = prompt.Field("adb path", adb, "", config.Vars.Binary.ADB)
	automator = prompt.Field("automator path", automator, "", config.Vars.Binary.Automator)

	if name == "" {
		return errors.New("missing: name")
	}

	commonOptions := commands.ScreenShotOptions{
		Name:      name,
		OutputDir: ".",
	}
	if platform == "android" {
		options := commands.ScreenShotAndroidOptions{
			ScreenShotOptions: commonOptions,
			ADB:               adb,
		}
		_, err = commands.ScreenShotAndroid(options)
	} else if platform == "ios" || platform == "ios-simulator" {
		options := commands.ScreenShotIOSOptions{
			ScreenShotOptions: commonOptions,
			Automator:         automator,
			Simulator:         platform == "ios-simulator",
		}
		err := commands.ScreenShotIOSPrepare(options)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
		_, err = commands.ScreenShotIOS(options)
	} else if platform == "desktop" {
		options := commands.ScreenShotDesktopOptions{
			ScreenShotOptions: commonOptions,
		}
		_, err = commands.ScreenShotDesktop(options)
	}
	return err

}
