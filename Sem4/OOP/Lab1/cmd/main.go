package main

import (
	"Lab1/cmd/cmdManager"
	"Lab1/services"
)

func main() {

	//Init hardware section
	var disks []services.Disk
	var rams []services.RAM
	var cpus []services.CPU

	backupManager := services.NewBackupManager(3, true)
	alertManager := services.NewAlertManager()

	screenManager := cmdManager.NewScreenManager()
	serverManager := cmdManager.NewServerManager()

	coolingCPUManager := services.NewCoolingManager(120)
	coolingRAMManager := services.NewCoolingManager(60)
	coolingDISKManager := services.NewCoolingManager(70)

	for {
		if !screenManager.IsScreen() {
			return
		}

		choice, _ := screenManager.PrintChoiceMenu()
		switch choice {

		case "1":
			screenManager.PrintHardwareMenu(&disks, &rams, &cpus, serverManager)
		case "2":
			screenManager.Output("Disks backup  turned on")
			serverManager.AddTask(func() {
				err := backupManager.StartPeriod(disks)
				alertManager.AddAlert(err)
			})
		case "3":
			screenManager.Output("Cooling monitor turn on")
			coolingCPUManager.Start()
			coolingRAMManager.Start()
			coolingDISKManager.Start()
		case "q":
			screenManager.Stop()
		}

	}
}
