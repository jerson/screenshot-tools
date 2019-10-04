package utils

import (
	"fmt"
	"os"
	"os/user"
)

// ExistsFile ...
func ExistsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// DesktopDir ...
func DesktopDir() (string, error) {

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", usr.HomeDir, "Desktop"), nil
}
