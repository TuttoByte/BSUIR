package services

import "errors"

type AccessControl struct {
	users map[uint64]bool
}

func NewAccessControl() *AccessControl {
	return &AccessControl{
		users: make(map[uint64]bool),
	}
}

func (ac *AccessControl) AddUser(user uint64, permission bool) error {

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
