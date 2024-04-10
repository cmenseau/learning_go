package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type errKind int

const (
	tooShort errKind = iota
	tooLong
	missingSpecialChars
	missingCapitalLetters
)

type PasswordError struct {
	kind errKind
}

func (p PasswordError) Error() string {
	switch p.kind {
	case tooShort:
		return "password must be at least 8 characters long"
	case tooLong:
		return "password must be at most 32 characters long"
	case missingSpecialChars:
		return "password must contain special characters"
	case missingCapitalLetters:
		return "password must contain capital letters"
	default:
		return ""
	}
}

var (
	TooShort              PasswordError = PasswordError{kind: tooShort}
	TooLong               PasswordError = PasswordError{kind: tooLong}
	MissingCapitalLetters PasswordError = PasswordError{kind: missingCapitalLetters}
	MissingSpecialChars   PasswordError = PasswordError{kind: missingSpecialChars}
)

func error_prototype() {

	var validatePwd = func(password string) error {
		if len(password) < 8 {
			return TooShort
		}
		if len(password) > 32 {
			return TooLong
		}
		if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return MissingCapitalLetters
		}
		if !strings.ContainsAny(password, "!:;,&*%$^") { //etc
			return MissingSpecialChars
		}
		return nil
	}

	fmt.Printf("%T : %[1]s\n", validatePwd("short"))
	fmt.Printf("%T : %[1]s\n", validatePwd("veryveryveryveryveryveryveryverylong"))
	fmt.Printf("%T : %[1]s\n", validatePwd("nocapitalletter"))
	fmt.Printf("%T : %[1]s\n", validatePwd("noSpecialChar"))
	fmt.Printf("%T : %[1]s\n", validatePwd("okayPassword!"))
}

func wrapping_errors() {

	var readFromFile = func() (string, error) {
		data, err := os.ReadFile("wrong.txt")
		if err != nil {
			err = fmt.Errorf("read from file : %w", err)
			return "", err
		}
		return string(data), nil
	}

	var getConfig = func() (string, error) {
		data, err := readFromFile()
		if err != nil {
			err = fmt.Errorf("get config : %w", err)
			return "", err
		}
		res := data[strings.Index(data, "config="):]
		return res, nil
	}

	config, err := getConfig()

	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Is(err, os.ErrNotExist))   // true
		fmt.Println(errors.Is(err, os.ErrPermission)) // false

		var e *os.PathError
		if errors.As(err, &e) {
			fmt.Println(e)
		}

		fmt.Println("UNWRAPPING")
		for err = errors.Unwrap(err); err != nil; err = errors.Unwrap(err) {
			fmt.Println(err)
		}
	} else {
		fmt.Println(config)
	}
}
func willPanic(str string) {
	panic("bad thing happened : " + str)
}

func panicking_recovering() {

	defer func() {
		if d := recover(); d != nil {
			fmt.Println("Recovered :", d)
		}
	}()
	willPanic("muhahhaha")
	fmt.Println("after")
}

func wontPanic(str string) (ret string, err error) {
	ret, err = "res", nil
	defer func() {
		if d := recover(); d != nil {
			fmt.Println("Recovered :", d)
			ret = "reco"
			return
		}
	}()
	willPanic(str)
	return
}

func main() {
	//error_prototype()
	//wrapping_errors()
	//panicking_recovering()
	fmt.Println(wontPanic("input"))
}
