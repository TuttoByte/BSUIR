package main

import (
	"Lab1/cmd/cmdManager"
)

func main() {

	//Init hardware section
	screenManager := cmdManager.NewScreenManager()

	for {
		if !screenManager.IsScreen() {
			return
		}

		choice, _ := screenManager.PrintChoiceMenu()
		switch choice {

		case "1":
			screenManager.PrintHardwareMenu()
		case "2":
			screenManager.SM.AddTask(func() {

				var err error
				if screenManager.SM.BackupManager.GetStatus() {
					err = screenManager.SM.BackupManager.StopPeriod()
					return
				}
				err = screenManager.SM.BackupManager.StartPeriod(screenManager.SM.GetDevices())
				if err != nil {
					screenManager.SM.AlertManager.AddAlert(err)
				}
			})

		case "q":
			screenManager.Stop()
		}

	}
}
