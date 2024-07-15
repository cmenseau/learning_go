package trip

import (
	"github.com/pkg/errors"
)

var session *UserSession
var dao_err = errors.New("TripDAO should not be invoked on an unit test.")
var userSession_err = errors.New("UserSession.GetLoggedUser() should not be called in an unit test")

type Trip struct {
}

type Service struct {
	tripDAO *Dao
}

func (service *Service) GetTripByUser(user *User) ([]Trip, error) {
	var trips []Trip
	friends, err := user.Friends()
	if err != nil {
		return trips, err
	}
	loggedUser, err := session.GetLoggedUser()
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
		return trips, errors.New("user not logged in")
	}
}

type UserSession struct {
}

func (userSession *UserSession) GetLoggedUser() (*User, error) {
	return nil, userSession_err
}
}

type User struct {
}

func (user *User) Friends() ([]User, error) {
	var friends []User
	return friends, nil
}

type Dao struct {
}

func (dao *Dao) FindTripByUser(user *User) ([]Trip, error) {
	return nil, dao_err
}
