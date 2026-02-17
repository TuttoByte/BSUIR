package services

import (
	"Lab1/services/network/control"
	"errors"
)

type AccessControl struct {
	usersPermisions map[uint64]control.Client
}

func NewAccessControl() *AccessControl {
	return &AccessControl{
		usersPermisions: make(map[uint64]control.Client),
	}
}

func (ac *AccessControl) AddUser(client control.Client) error {

	_, ok := ac.users[user]
	if !ok {
		ac.users[user] = permission
	} else {
		return errors.New("User already exists")
	}
	return nil
}

func (ac *AccessControl) RemoveUser(user uint64) error {
	_, ok := ac.users[user]
	if !ok {
		return errors.New("User does not exist")
	}
	delete(ac.users, user)
	return nil
}

func (ac *AccessControl) SetUserPermission(user uint64, permission bool) error {
	_, ok := ac.users[user]
	if !ok {
		return errors.New("User does not exist")
	}
	ac.users[user] = permission
	return nil
}
