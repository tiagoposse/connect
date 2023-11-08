package utils

import (
	"errors"
	"fmt"

	"github.com/go-faster/jx"
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

// RawError renders err as json string.
func RawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}
