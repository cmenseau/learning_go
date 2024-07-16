package trip

import (
	"github.com/pkg/errors"
)

var loggedUserGetter LoggedUserGetter = &UserSession{}
var dao_err = errors.New("TripDAO should not be invoked on an unit test.")
var userSession_err = errors.New("UserSession.GetLoggedUser() should not be called in an unit test")
var notLoggedUser_err = errors.New("user not logged in")

type Trip struct {
}

type Service struct {
	tripDAO DaoTripFinder
}

func (service *Service) GetTripByUser(user Extrovert) ([]Trip, error) {
	var trips []Trip

	loggedUser, err := loggedUserGetter.GetLoggedUser()

	if err != nil {
		return trips, err
	}

	if loggedUser == nil {
		return trips, notLoggedUser_err
	}

	if loggedUser.IsFriendWith(user) {
		return service.tripDAO.FindTripByUser(user)
	}

	return trips, err
}

type LoggedUserGetter interface {
	GetLoggedUser() (*User, error)
}

type UserSession struct {
}

func (userSession *UserSession) GetLoggedUser() (*User, error) {
	return nil, userSession_err
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
	friends, err := extrovert.GetFriends()

	if err != nil {
		return false
	}

	for _, friend := range friends {
		if *user == friend {
			return true
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
	return nil, dao_err
}
