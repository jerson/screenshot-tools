package utils

import "os"

// ExistsFile ...
func ExistsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
