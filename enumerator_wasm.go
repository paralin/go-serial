//go:build wasip1
// +build wasip1

package serial

import (
	"errors"
)

func nativeGetPortsList() ([]string, error) {
	return nil, errors.New("nativeGetPortsList is not supported on GOARCH=wasm")
}
