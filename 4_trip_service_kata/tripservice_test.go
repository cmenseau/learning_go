package trip

import (
	"testing"

	"github.com/pkg/errors"
)

func TestEmptyUser(test *testing.T) {

	var user User

	var friends, err = user.Friends()

	if friends != nil || err != nil {
		test.Errorf("expected friends, err to be (%v, %v) got (%v, %v)", nil, nil, friends, err)
	}
}

func TestDao(test *testing.T) {

	var dao Dao
	var user *User

	var trips, err = dao.FindTripByUser(user)

	expected_err := dao_err

	if trips != nil || !errors.Is(err, expected_err) {
		test.Errorf("expected friends, err to be (%v, %v) got (%v, %v)",
			nil, expected_err, trips, err)
	}
}

func TestUserSession(test *testing.T) {

	var userSession UserSession

	var user, err = userSession.GetLoggedUser()

	expected_err := userSession_err

	if user != nil || !errors.Is(err, expected_err) {
		test.Errorf("expected friends, err to be (%v, %v) got (%v, %v)",
			nil, expected_err, user, err)
	}
}

