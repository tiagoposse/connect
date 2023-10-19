package utils

import (
	"errors"
	"fmt"
)

func NewMaxItemsError(msg string) maxItemsError {
	return maxItemsError{msg: msg}
}

type maxItemsError struct {
	msg string
}

func (m maxItemsError) Error() string {
	return fmt.Sprintf("Max number reached for %s", m.msg)
}

func IsMaxItemsError(err error) bool {
	return errors.As(err, &maxItemsError{})
}
