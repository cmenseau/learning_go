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

func (service *Service) GetTripByUser(user Friender) ([]Trip, error) {
	var trips []Trip
	friends, err := user.Friends()
	if err != nil {
		return trips, err
	}
	loggedUser, err := loggedUserGetter.GetLoggedUser()
	if err != nil {
		return trips, err
	}
	var isFriend bool
	if loggedUser != nil {
		for _, friend := range friends {
			if *loggedUser == friend {
				isFriend = true
				break
			}
		}
		if isFriend {
			return service.tripDAO.FindTripByUser(user)
		}
		return trips, err
	} else {
		return trips, notLoggedUser_err
	}
}

type LoggedUserGetter interface {
	GetLoggedUser() (*User, error)
}

type UserSession struct {
}

func (userSession *UserSession) GetLoggedUser() (*User, error) {
	return nil, userSession_err
}

type Friender interface {
	Friends() ([]User, error)
}

type User struct {
}

func (user *User) Friends() ([]User, error) {
	var friends []User
	return friends, nil
}

type DaoTripFinder interface {
	FindTripByUser(user Friender) ([]Trip, error)
}

type Dao struct {
}

func (dao *Dao) FindTripByUser(user Friender) ([]Trip, error) {
	return nil, dao_err
}
