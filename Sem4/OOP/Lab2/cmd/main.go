package main

import (
	"Lab1/cmd/cmdManager"
	"Lab1/services"
	"Lab1/services/devices"
	"Lab1/services/managres"
)

func main() {

	//Init hardware section
	var disks []devices.Disk
	var rams []devices.RAM
	var cpus []devices.CPU

	backupManager := managres.NewBackupManager(3, true)
	alertManager := services.NewAlertManager()

	screenManager := cmdManager.NewScreenManager()
	serverManager := cmdManager.NewServerManager()

	coolingCPUManager := managres.NewCoolingManager(120)
	coolingRAMManager := managres.NewCoolingManager(60)
	coolingDISKManager := managres.NewCoolingManager(70)

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
