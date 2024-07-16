package trip

import (
	"errors"
	"testing"
)

func TestEmptyUser(test *testing.T) {

	var user User

	var friends, err = user.GetFriends()

	if friends != nil || err != nil {
		test.Errorf("expected friends, err to be (%v, %v) got (%v, %v)", nil, nil, friends, err)
	}
}

func TestDao(test *testing.T) {

	var dao Dao
	var user *User

	var trips, err = dao.FindTripByUser(user)

	expected_err := errDao

	if trips != nil || !errors.Is(err, expected_err) {
		test.Errorf("expected friends, err to be (%v, %v) got (%v, %v)",
			nil, expected_err, trips, err)
	}
}

func TestUserSession(test *testing.T) {

	var userSession UserSession

	var user, err = userSession.GetLoggedUser()

	expected_err := errUserSession

	if user != nil || !errors.Is(err, expected_err) {
		test.Errorf("expected friends, err to be (%v, %v) got (%v, %v)",
			nil, expected_err, user, err)
	}
}

type MockUser struct {
	getFriends func() ([]User, error) // set it to control behavior
}

// implement Friender
func (mu MockUser) GetFriends() ([]User, error) {
	return mu.getFriends()
}

func TestFriendsError(test *testing.T) {

	var svc Service = Service{tripDAO: &Dao{}}
	var mu MockUser = MockUser{
		getFriends: func() ([]User, error) {
			return nil, errors.New("Some error")
		},
	}

	trips, err := svc.GetTripByUser(mu)

	if trips != nil || err == nil {
		test.Errorf("expected trips, err to be (%v, %v) got (%v, %v)",
			nil, "Some error", trips, err)
	}
}

type MockGetLoggedUser struct {
	getLoggedUser func() (*User, error) // set it to control behavior
}

func (mock *MockGetLoggedUser) GetLoggedUser() (*User, error) {
	return mock.getLoggedUser()
}

func TestNoLoggedUser(test *testing.T) {

	loggedUserGetter = &MockGetLoggedUser{
		getLoggedUser: func() (*User, error) {
			return nil, nil
		},
	}

	var svc = &Service{
		tripDAO: &Dao{},
	}

	var user User

	trips, err := svc.GetTripByUser(&user)

	if trips != nil || err != errNotLoggedUser {
		test.Errorf("expected trips, err to be (%v, %v) got (%v, %v)",
			nil, errNotLoggedUser, trips, err)
	}
}

func TestLoggedUser(test *testing.T) {

	loggedUserGetter = &MockGetLoggedUser{
		getLoggedUser: func() (*User, error) {
			return &User{}, nil
		},
	}

	var svc = &Service{
		tripDAO: &Dao{},
	}

	var user User

	trips, err := svc.GetTripByUser(&user)

	if trips != nil || err != nil {
		test.Errorf("expected trips, err to be (%v, %v) got (%v, %v)",
			nil, nil, trips, err)
	}
}

func TestLoggedUserError(test *testing.T) {

	var loggedUserErr = errors.New("Some Error")

	loggedUserGetter = &MockGetLoggedUser{
		getLoggedUser: func() (*User, error) {
			return &User{}, loggedUserErr
		},
	}

	var svc = &Service{
		tripDAO: &Dao{},
	}

	var user User

	trips, err := svc.GetTripByUser(&user)

	if trips != nil || err != loggedUserErr {
		test.Errorf("expected trips, err to be (%v, %v) got (%v, %v)",
			nil, loggedUserErr, trips, err)
	}
}

func TestGetTripOfStranger(test *testing.T) {

	var mu MockUser = MockUser{
		getFriends: func() ([]User, error) {
			return []User{User{}}, nil
		},
	}

	loggedUserGetter = &MockGetLoggedUser{
		getLoggedUser: func() (*User, error) {
			return &User{}, nil
		},
	}

	var svc = &Service{
		tripDAO: &Dao{},
	}

	trips, err := svc.GetTripByUser(mu)

	if trips != nil || !errors.Is(err, errDao) {
		test.Errorf("expected trips, err to be (%v, %v) got (%v, %v)",
			nil, errDao, trips, errDao)
	}
}

type MockDao struct {
	findTrip func() ([]Trip, error) // set it to control behavior
}

func (dao *MockDao) FindTripByUser(user Extrovert) ([]Trip, error) {
	return dao.findTrip()
}

func TestGetTripOfFriend(test *testing.T) {

	var friend User = User{}

	var mu MockUser = MockUser{
		getFriends: func() ([]User, error) {
			return []User{User{}, friend}, nil
		},
	}
	loggedUserGetter = &MockGetLoggedUser{
		getLoggedUser: func() (*User, error) {
			return &User{}, nil
		},
	}

	var mockDao MockDao = MockDao{
		findTrip: func() ([]Trip, error) {
			return []Trip{Trip{}, Trip{}}, nil
		},
	}

	var svc = &Service{
		tripDAO: &mockDao,
	}

	trips, err := svc.GetTripByUser(mu)

	if len(trips) != 2 || err != nil {
		test.Errorf("expected trips, err to be (%v, %v) got (%v, %v)",
			nil, errDao, trips, errDao)
	}
}
