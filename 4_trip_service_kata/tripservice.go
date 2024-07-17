package trip

import (
	"errors"
)

var loggedUserGetter LoggedUserGetter = &UserSession{}
var errDao = errors.New("TripDAO should not be invoked on an unit test")
var errUserSession = errors.New("UserSession.GetLoggedUser() should not be called in an unit test")
var errNotLoggedUser = errors.New("user not logged in")

type Trip struct {
}

type Service struct {
	tripFinder DaoTripFinder
}

func (service *Service) GetTripByUser(user Extrovert) ([]Trip, error) {
	loggedUser, err := service.GetCurrentLoggedUser()

	if err != nil {
		return nil, err
	}

	if loggedUser.IsFriendWith(user) {
		return service.tripFinder.FindTripByUser(user)
	}

	return nil, nil
}

func (service *Service) GetCurrentLoggedUser() (*User, error) {
	loggedUser, err := loggedUserGetter.GetLoggedUser()

	if err != nil {
		return nil, err
	}

	if loggedUser == nil {
		return nil, errNotLoggedUser
	}

	return loggedUser, nil
}

type LoggedUserGetter interface {
	GetLoggedUser() (*User, error)
}

type UserSession struct {
}

func (userSession *UserSession) GetLoggedUser() (*User, error) {
	return nil, errUserSession
}

type Extrovert interface {
	GetFriends() ([]User, error)
}

type User struct {
}

func (user *User) GetFriends() ([]User, error) {
	var friends []User
	return friends, nil
}

func (user *User) IsFriendWith(extrovert Extrovert) bool {

	if friends, err := extrovert.GetFriends(); err == nil {

		for _, friend := range friends {
			if *user == friend {
				return true
			}
		}
	}
	return false
}

type DaoTripFinder interface {
	FindTripByUser(user Extrovert) ([]Trip, error)
}

type Dao struct {
}

func (dao *Dao) FindTripByUser(user Extrovert) ([]Trip, error) {
	return nil, errDao
}
