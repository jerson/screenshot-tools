package terminal

import (
	"os/exec"
	"runtime"
)

// RestoreInput ...
func RestoreInput() error {
	err := exec.Command("stty", "sane").Run()
	if err != nil {
		return err
	}
	return nil
}

// HideInput ...
func HideInput() error {
	if runtime.GOOS == "linux" {
		err := exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		if err != nil {
			return err
		}
	}
	if runtime.GOOS == "darwin" {
		err := exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// InputWithoutBreakLine ...
func InputWithoutBreakLine() error {
	if runtime.GOOS == "linux" {
		err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		if err != nil {
			return err
		}
	}
	if runtime.GOOS == "darwin" {
		err := exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
		if err != nil {
			return err
		}
	}
	return nil
}
