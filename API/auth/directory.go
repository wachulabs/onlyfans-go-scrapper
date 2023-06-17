package auth

import (
	"github.com/pkg/errors"
	"os"
)

// ConfirmDirectoryExistence confirms whether a directory exits
// It returns a boolean value and a error
func ConfirmDirectoryExistence(directory ...string) (bool, error) {
	var path string
	var err error
	if len(directory) == 0 {
		e := errors.New("you must input a directory to check")
		err = errors.Wrap(e, "Directory does not exits")
		return false, err
	} else {
		for _, val := range directory {
			path += val
		}
	}
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {

		return false, nil
	}
	return false, err
}

// CreateDirectory creates directory if directory does not exist
func CreateDirectory(directory ...string) error {
	var err error
	var path string
	for _, val := range directory {
		path += val
	}
	bl, _ := ConfirmDirectoryExistence(path)
	if bl {
		e := errors.New("Directory exists")
		return errors.Wrap(e, "Can not create an existing directory")
	}

	if len(directory) == 0 {
		err = os.Mkdir("unnamedfolder", os.ModePerm)
		if err != nil {
			return err
		} else {
			return errors.New("Creating unnamed folder. No name was given")
		}
	} else {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}
