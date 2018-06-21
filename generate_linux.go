package uuidgen

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

func generate() ([]byte, error) {
	b := make([]byte, 16)

	n, err := unix.Getrandom(b, 0)
	if err != nil {
		return nil, errors.Wrap(err, "Received error attempting syscall")
	}

	if n != len(b) {
		return nil, errors.New("Syscall did not generate sufficient bytes")
	}

	return b, nil
}
