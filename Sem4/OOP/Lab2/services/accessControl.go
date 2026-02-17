package services

import (
	"Lab1/services/network/control"
	"errors"
)

type AccessControl struct {
	usersPermisions map[uint64]*control.Client
}

func NewAccessControl() *AccessControl {
	return &AccessControl{
		usersPermisions: make(map[uint64]*control.Client),
	}
}

func (ac *AccessControl) AddUser(client control.Client) error {

	_, ok := ac.usersPermisions[client.GetID()]
	if !ok {
		ac.usersPermisions[client.GetID()] = &client
	} else {
		return errors.New("User already exists")
	}
	return nil
}

func (ac *AccessControl) RemoveUser(client control.Client) error {
	_, ok := ac.usersPermisions[client.GetID()]
	if !ok {
		return errors.New("User does not exist")
	}
	delete(ac.usersPermisions, client.GetID())
	return nil
}

func (ac *AccessControl) SetUserPermission(id uint64, permission bool) error {
	_, ok := ac.usersPermisions[id]
	if !ok {
		return errors.New("User does not exist")
	}
	ac.usersPermisions[id].SetControlStatus(permission)
	return nil
}
