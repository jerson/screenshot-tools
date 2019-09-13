package prompt

import (
	"errors"
	"runtime"

	ui "github.com/VladimirMarkelov/clui"
	log "github.com/sirupsen/logrus"
	"gopkg.in/AlecAivazis/survey.v1"
)

func chooseDir(output chan string, title, input string, closeOnFinish bool) {

	go func() {
		ui.InitLibrary()
		defer ui.DeinitLibrary()
		dialog := ui.CreateFileSelectDialog(title, "", input, true, true)
		dialog.OnClose(func() {
			selected := dialog.Selected
			path := dialog.FilePath
			if selected && path == "" {
				path = input
			}
			defer func() {
				ui.WindowManager().DestroyWindow(dialog.View)
				ui.WindowManager().BeginUpdate()
				ui.WindowManager().EndUpdate()
				if closeOnFinish && runtime.GOOS != "windows" {
					ui.DeinitLibrary()
				}
			}()

			output <- path

		})
		ui.MainLoop()
	}()
}

func chooseFile(output chan string, title, input, fileMasks string, closeOnFinish bool) {

	go func() {
		ui.InitLibrary()
		defer ui.DeinitLibrary()
		dialog := ui.CreateFileSelectDialog(title, fileMasks, input, false, true)
		dialog.OnClose(func() {
			selected := dialog.Selected
			path := dialog.FilePath
			if selected && path == "" {
				path = input
			}
			defer func() {
				ui.WindowManager().DestroyWindow(dialog.View)
				ui.WindowManager().BeginUpdate()
				ui.WindowManager().EndUpdate()
				if closeOnFinish && runtime.GOOS != "windows" {
					ui.DeinitLibrary()
				}
			}()
			output <- path
		})
		ui.MainLoop()
	}()
}
func File(name, value, fileMasks, defaultValue string, closeOnFinish bool) string {
	if value == "" {
		output := make(chan string)
		chooseFile(output, name, defaultValue, fileMasks, closeOnFinish)
		<-output
		value = <-output
	}
	return value
}
func Dir(name, value, defaultValue string, closeOnFinish bool) string {
	if value == "" {
		output := make(chan string)
		chooseDir(output, name, defaultValue, closeOnFinish)
		<-output
		value = <-output
	}
	return value
}

func Field(name, value, help, defaultValue string) string {
	if value == "" {
		prompt := &survey.Input{
			Message: name,
			Default: defaultValue,
			Help:    help,
		}
		err := survey.AskOne(prompt, &value, requiredField)
		if err != nil {
			log.Warn(err)
			return value
		}
	}
	return value
}

func PasswordField(name, value, help, defaultValue string) string {
	if value == "" {
		value = defaultValue
		prompt := &survey.Password{
			Message: name,
			Help:    help,
		}
		err := survey.AskOne(prompt, &value, requiredField)
		if err != nil {
			log.Warn(err)
			return value
		}

	}
	return value
}

func requiredField(ans interface{}) error {
	input := ans.(string)
	if len(input) < 1 {
		return errors.New("required field")
	}
	return nil
}
