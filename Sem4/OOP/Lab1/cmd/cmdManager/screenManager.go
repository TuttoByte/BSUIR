package cmdManager

import (
	"errors"
	"fmt"
	"github.com/inancgumus/screen"
	"os"
)

type ScreenManager struct {
	screen bool
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
