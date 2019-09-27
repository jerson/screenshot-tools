package main

//go:generate rm -rf assets/automator.zip
//go:generate zip -r assets/automator.zip assets/automator

import (
	"os"
	"runtime"

	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"screenshot-tools/actions"
	"screenshot-tools/modules/config"
	"screenshot-tools/modules/osx"
)

func setup() {
	log.SetLevel(log.DebugLevel)
	if runtime.GOOS == "windows" {
		log.SetFormatter(&log.TextFormatter{
			DisableTimestamp: true,
			ForceColors:      false,
			QuoteEmptyFields: false,
		})
	}

	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}

	box := packr.New("Assets", "./assets")
	err = osx.LoadAssets(box)
	if err != nil && runtime.GOOS == "darwin" {
		panic(err)
	}

}
func main() {

	setup()

	app := cli.NewApp()
	app.Name = "ScreenShot Tools"
	app.Version = "0.1.10"
	app.Usage = ""

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:        "dump-config",
			Aliases:     []string{"dc"},
			Flags:       []cli.Flag{},
			Category:    "debug",
			Description: "Dump sample config",
			Usage:       "dump-config",
			UsageText: `
dump-config`,
			Action: actions.DumpConfig,
		},
		{
			Name:    "merge-images",
			Aliases: []string{"m"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Value: config.Vars.Dir.Input,
					Usage: "Input dir",
				},
				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: "Output file",
				},
			},
			Category:    "evidences",
			Description: "merge images into one merged file",
			Usage:       "merge-images",
			UsageText: `
merge-images
merge-images -o ./output.png
merge-images --input=./images
merge-images -i ./images
merge-images --input=./images --output=./output.png
merge-images -i ./images -o ./output.png`,
			Action: actions.MergeImages,
		},
		{
			Name:    "compress",
			Aliases: []string{"c"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Value: config.Vars.Dir.Input,
					Usage: "Input dir",
				},
				cli.StringFlag{
					Name:  "output, o",
					Value: config.Vars.Dir.Output,
					Usage: "Output dir",
				},
			},
			Category:    "evidences",
			Description: "compress images grouped by report",
			Usage:       "compress",
			UsageText: `
compress
compress -o ./output
compress --input=./images
compress -i ./images
compress --input=./images --output=./output
compress -i ./images -o ./output`,
			Action: actions.Compress,
		},
		{
			Name:    "report",
			Aliases: []string{"r"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input, i",
					Value: config.Vars.Dir.Input,
					Usage: "Input dir",
				},
			},
			Category:    "debug",
			Description: "show report for debug purposes",
			Usage:       "report",
			UsageText: `
report
report --input=./images
report -i ./images`,
			Action: actions.Report,
		},
		{
			Name:    "screenshot-session",
			Aliases: []string{"ss"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "adb, a",
					Value: config.Vars.Binary.ADB,
					Usage: "ADB path used when platform=android",
				},
				cli.StringFlag{
					Name:  "automator, au",
					Value: config.Vars.Binary.Automator,
					Usage: "Automator used when platform=ios",
				},
				cli.StringFlag{
					Name:  "platform, p",
					Value: "android",
					Usage: "Platform: ios,android,ios-simulator,desktop",
				},
			},
			Category:    "screenshot",
			Description: "start session for take many screenshots",
			Usage:       "screenshot-session",
			UsageText: `
screenshot-session
screenshot-session name`,
			Action: actions.ScreenShotSession,
		},
		{
			Name:    "screenshot",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "adb, a",
					Value: config.Vars.Binary.ADB,
					Usage: "ADB path used when platform=android",
				},
				cli.StringFlag{
					Name:  "automator, au",
					Value: config.Vars.Binary.Automator,
					Usage: "Automator used when platform=ios",
				},
				cli.StringFlag{
					Name:  "platform, p",
					Value: "android",
					Usage: "Platform: ios,android,ios-simulator,desktop",
				},
			},
			Category:    "screenshot",
			Description: "capture screenshot",
			Usage:       "screenshot",
			UsageText: `
screenshot
screenshot "sample name"`,
			Action: actions.ScreenShot,
		},
		{
			Name:    "upload-nexus",
			Aliases: []string{"un"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "username, u",
					Value: config.Vars.Nexus.Username,
					Usage: "Nexus username",
				},
				cli.StringFlag{
					Name:  "password, p",
					Value: config.Vars.Nexus.Password,
					Usage: "Nexus password",
				},
				cli.StringFlag{
					Name:  "project, pr",
					Value: config.Vars.Nexus.Project,
					Usage: "Nexus project",
				},
				cli.StringFlag{
					Name:  "server, s",
					Value: config.Vars.Nexus.Server,
					Usage: "Nexus server template",
				},
				cli.StringFlag{
					Name:  "name, n",
					Value: "",
					Usage: "Filename with extension",
				},
			},
			Category:    "nexus",
			Description: "upload android or ios binaries to nexus",
			Usage:       "upload-nexus app-10-10-2010.apk",
			UsageText: `
upload-nexus app-10-10-2010.apk
upload-nexus app-10-10-2010.ipa
upload-nexus app-10-10-2010.zip`,
			Action: actions.UploadNexus,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
