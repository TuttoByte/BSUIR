package cmdManager

import (
	"Lab1/services/devices"
	"errors"
	"fmt"
	"github.com/inancgumus/screen"
	"os"
)

type ScreenManager struct {
	screen bool
	sm     *ServerManager
}

func NewScreenManager() *ScreenManager {
	err := getStartedMassage()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
	return &ScreenManager{
		screen: true,
	}
}

func (s *ScreenManager) Start() {
	s.screen = true
}

func (s *ScreenManager) Stop() {
	s.screen = false
}

func (s *ScreenManager) IsScreen() bool {
	return s.screen
}

func (s *ScreenManager) Output(text string) {
	fmt.Println(text)
	_, _ = fmt.Scanln()
}
func getStartedMassage() error {
	fmt.Println("Hello, this is simpe data centre controll")
	fmt.Println("Print 1 to exit")
	fmt.Println("Print any button to start")
	var choice string
	_, _ = fmt.Scanln(&choice)

	if choice == "1" {
		return errors.New("Stop")
	}
	return nil

}

func (s *ScreenManager) PrintChoiceMenu() (string, error) {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("To configurate hardware press 1")
	fmt.Println("To configurate software press 2")
	fmt.Println("To exit pree q")

	var choice string
	_, _ = fmt.Scanln(&choice)

	return choice, nil
}

func (s *ScreenManager) PrintHardwareMenu(dicks *[]devices.Disk, rams *[]devices.RAM, cpus *[]devices.CPU, sm *ServerManager) {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("Hardware controll section")
	fmt.Println("Print 1 to control RAM")
	fmt.Println("Print 2 to control CPU")
	fmt.Println("Print 3 to control DISK")
	fmt.Println("Print b to get back")
	var choice string
	_, _ = fmt.Scanln(&choice)

	s.ShowHardwareConfMenu(choice, sm)

}

func (s *ScreenManager) ShowHardwareConfMenu(choice string, sm *ServerManager) {
	screen.Clear()
	screen.MoveTopLeft()

	switch choice {
	case "1":
		s.ramHardwareMenu(rams, "RAM", sm)
	case "2":
		s.cpuHardwareMenu(cpus, "CPU", sm)
	case "3":
		s.diskHardwareMenu(dicks, "DISKS", sm)

	}
}
func (s *ScreenManager) ramHardwareMenu(storage *[]devices.RAM, htype string, sm *ServerManager) string {
	choice := printIn(htype)

	switch choice {

	case "1":
		sm.AddTask(func() {
			ram := devices.NewRAM(64, 5200)
			*storage = append(*storage, *ram)
		})
	case "2":

	case "3":
		showList[devices.RAM](*storage)
		_, _ = fmt.Scanln()

	}

	return choice
}

func (s *ScreenManager) cpuHardwareMenu(storage *[]devices.CPU, htype string, sm *ServerManager) string {
	choice := printIn(htype)

	switch choice {

	case "1":
		sm.AddTask(func() {
			ram := devices.NewCPU(48, 96, 4, "INTEL")
			*storage = append(*storage, *ram)
		})
	case "2":

	case "3":
		showList[devices.CPU](*storage)
		_, _ = fmt.Scanln()

	}

	return choice
}

func (s *ScreenManager) diskHardwareMenu(storage *[]devices.Disk, htype string, sm *ServerManager) string {
	choice := printIn(htype)
	fmt.Println(storage)

	switch choice {

	case "1":
		sm.AddTask(func() {
			ram := devices.NewDisk("HDD", 4096)
			*storage = append(*storage, *ram)
		})

	case "2":

	case "3":
		showList[devices.Disk](*storage)
		_, _ = fmt.Scanln()

	}

	return choice
}

func printIn(htype string) string {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("Hello, this is ", htype, "controll")
	fmt.Println("Print 1 to add", htype)
	fmt.Println("Print 2 to delete", htype)
	fmt.Println("Print 3 to list all", htype)
	fmt.Println("Print b to get back")
	var choice string
	_, _ = fmt.Scanln(&choice)
	return choice
}

func showList[T any](list []T) {
	for _, item := range list {
		fmt.Println(item)
	}
}
