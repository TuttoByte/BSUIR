package main

import (
	"Lab1/cmd/cmdManager"
)

func main() {

	screenManager := cmdManager.NewScreenManager()
	_ = cmdManager.NewServerManager()

	for {
		if !screenManager.IsScreen() {
			return
		}

		choice, _ := screenManager.PrintChoiceMenu()
		switch choice {

		case "q":
			screenManager.Stop()
		}

	}
}
