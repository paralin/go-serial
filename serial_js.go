//go:build js
// +build js

package serial

import (
	"errors"
)

func nativeOpen(portName string, mode *Mode) (Port, error) {
	return nil, errors.New("nativeOpen is not supported on GOOS=js")
}
