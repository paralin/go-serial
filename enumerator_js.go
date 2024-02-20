//go:build js
// +build js

package serial

import (
	"errors"
)

func nativeGetPortsList() ([]string, error) {
	return nil, errors.New("nativeGetPortsList is not supported on GOOS=js")
}
