package neterror

import (
	"errors"
	"net"
	"syscall"
)

// GetNetError recursively checks the chain of errors until it finds an error
// from the net package that fulfills the net.Error interface. If we find one
// that means its safe to say that a network error has occurred. See:
// https://godoc.org/net#Error for more information on the net.Error interface.
func GetNetError(err error) (net.Error, bool) {
	for {
		if err == nil {
			return nil, false
		}

		// Ignore syscall.Errno errors, which weirdly enough satisfy the
		// net.Error interface
		switch err.(type) {
		case syscall.Errno, *syscall.Errno:
			return nil, false
		}

		if netError, ok := err.(net.Error); ok {
			return netError, true
		}

		err = errors.Unwrap(err)
	}
}
