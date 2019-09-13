//Package config ...
package config

import (
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
)

// Resize ...
type Resize struct {
	Height  int `toml:"height" default:"600"`
	Columns int `toml:"columns" default:"2"`
}

// Dir ...
type Dir struct {
	Automator string `toml:"automator" default:".screenshot_tools"`
	Input     string `toml:"input" default:""`
	Output    string `toml:"output" default:"./output"`
}

// Nexus ...
type Nexus struct {
	Username string `toml:"username" default:""`
	Password string `toml:"password" default:""`
	Project  string `toml:"project" default:"app"`
	// project platform name
	Server string `toml:"server" default:"http://server/repository/%s-%s/builds/%s"`
}

// Binary ...
type Binary struct {
	ADB       string `toml:"adb" default:"adb"`
	XCode     string `toml:"xcode" default:"Xcode"`
	Automator string `toml:"automator" default:"/usr/bin/automator"`
}

//Vars ...
var Vars = struct {
	Debug  bool   `toml:"debug" default:"false"`
	Dir    Dir    `toml:"dir"`
	Nexus  Nexus  `toml:"nexus"`
	Binary Binary `toml:"binary"`
	Resize Resize `toml:"resize"`
}{}

//ReadDefault ...
func ReadDefault() error {
	file, err := filepath.Abs("./config.toml")
	if err != nil {
		return err
	}
	return Read(file)
}

//Read ...
func Read(file string) error {

	config := configor.New(&configor.Config{ENVPrefix: "ST", Debug: false, Verbose: false})
	if ExistsFile(file) {
		return config.Load(&Vars, file)
	}
	return config.Load(&Vars)
}

// ExistsFile ...
func ExistsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
