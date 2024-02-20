//go:build wasip1
// +build wasip1

package serial

import (
	"errors"
)

func nativeOpen(portName string, mode *Mode) (Port, error) {
	return nil, errors.New("nativeOpen is not supported on GOARCH=wasm")
}
