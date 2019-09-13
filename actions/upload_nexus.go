package actions

import (
	"errors"
	"path/filepath"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"

	"screenshot_tools/commands"
	"screenshot_tools/modules/config"
	"screenshot_tools/modules/prompt"
	"screenshot_tools/modules/utils"
)

// UploadNexus ...
func UploadNexus(c *cli.Context) error {

	username := c.String("username")
	password := c.String("password")
	server := c.String("server")
	project := c.String("project")
	name := c.String("name")
	file := c.Args().Get(0)

	file = prompt.File("Choose file to upload", file, "*.apk,*.ipa,*.zip", "", true)
	if file == "" {
		return errors.New("missing file")
	}

	log.Infof("Using file: %s", file)
	name = prompt.Field("Filename", name, "", filepath.Base(file))
	username = prompt.Field("Username", username, "", "")
	password = prompt.PasswordField("Password", password, "", "")
	project = prompt.Field("Project", project, "", "")
	server = prompt.Field("Server", server, "", config.Vars.Nexus.Server)

	if username == "" || password == "" || project == "" || server == "" {
		return errors.New("missing fields")
	}

	options := commands.UploadNexusOptions{
		UploadOptions: commands.UploadOptions{
			File: file,
			Name: name,
		},
		Username: username,
		Password: password,
		Project:  project,
		Server:   server,
	}
	url, err := commands.UploadNexus(options)
	if err != nil {
		return err
	}

	log.Info("Opening browser...")
	return utils.OpenBrowser(url)
}
