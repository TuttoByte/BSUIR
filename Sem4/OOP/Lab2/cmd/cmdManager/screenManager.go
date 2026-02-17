package cmdManager

import (
	"Lab1/interfaces"
	"Lab1/services/devices"
	"errors"
	"fmt"
	"github.com/inancgumus/screen"
	"os"
	"reflect"
)

type ScreenManager struct {
	screen bool
	SM     *ServerManager
}

func NewScreenManager() *ScreenManager {
	err := getStartedMassage()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
	return &ScreenManager{
		screen: true,
		SM:     NewServerManager(),
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
	if !s.SM.BackupManager.GetStatus() {
		fmt.Println("To Start monitoring system press 2")
	} else {
		fmt.Println("To Stop monitoring system press 2")
	}
	fmt.Println("To exit preess q")

	var choice string
	_, _ = fmt.Scanln(&choice)

	return choice, nil
}

func (s *ScreenManager) PrintHardwareMenu() {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("Hardware controll section")
	fmt.Println("Print 1 to control RAM")
	fmt.Println("Print 2 to control CPU")
	fmt.Println("Print 3 to control DISK")
	fmt.Println("Print b to get back")
	var choice string
	_, _ = fmt.Scanln(&choice)

	s.ShowHardwareConfMenu(choice)

}

func (s *ScreenManager) ShowHardwareConfMenu(choice string) {
	screen.Clear()
	screen.MoveTopLeft()

	var device interfaces.Device

	switch choice {
	case "1":
		device = &devices.RAM{}
	case "2":
		device = &devices.CPU{}
	case "3":
		device = &devices.Disk{}
	case "b":
		return
	}
	s.createHardwareMenu(device)

}
func (s *ScreenManager) createHardwareMenu(d interfaces.Device) {
	choice := printIn(d)

	switch choice {

	case "1":
		s.SM.AddTask(func() {
			newDevice(s.SM, d)
		})
	case "2":

	case "3":
		showList(s.SM.devices, d)
		_, _ = fmt.Scanln()
	case "b":
		return
	}

}

func printIn(d interfaces.Device) string {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("Hello, this is ", d.GetType(), "controll")
	fmt.Println("Print 1 to add", d.GetType())
	fmt.Println("Print 2 to delete", d.GetType())
	fmt.Println("Print 3 to list all", d.GetType())
	fmt.Println("Print b to get back")
	var choice string
	_, _ = fmt.Scanln(&choice)
	return choice
}

func showList(list []interfaces.Device, t interfaces.Device) {
	mainType := reflect.TypeOf(t)
	for _, item := range list {
		if reflect.TypeOf(item) == mainType {
			fmt.Println(item)
		}
	}
}

func newDevice(sm *ServerManager, dev interface{}) {

	screen.Clear()
	screen.MoveTopLeft()
	var d interfaces.Device

	switch v := (dev).(type) {

	case *devices.RAM:
		d = devices.NewRAM(64, 5200)
		fmt.Println("Device", v.GetType(), "successfully created")

	case *devices.CPU:
		d = devices.NewCPU(12, 16, 4, "Intel")
		fmt.Println("Device", v.GetType(), "successfully created")

	case *devices.Disk:
		d = devices.NewDisk("SSD", 1024)

		fmt.Println("Device", v.GetType(), "successfully created")

	default:
		fmt.Println("Device", v, "not recognized")

	}

	sm.devices = append(sm.devices, d)

}
