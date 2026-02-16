package services

import "fmt"

type AlertManager struct {
	alerts []error
}

func NewAlertManager() *AlertManager {
	return &AlertManager{
		alerts: make([]error, 0),
	}
}

func (a *AlertManager) GetAlerts() {
	for _, err := range a.alerts {
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (a *AlertManager) AddAlert(err error) {
	if err != nil {
		a.alerts = append(a.alerts, err)
	}
}
