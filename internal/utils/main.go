package utils

import "errors"

func CreateJsonReadError(err error) error {
	return errors.New("failed to understand JSON - " + err.Error())
}
