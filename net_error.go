package neterror

import (
	"errors"
	"net"
)

// GetNetError recursively checks the chain of errors until it finds an error
// that fulfills the net.Error interface. If we find one that means its safe to
// say that a network error has occurred. See: https://godoc.org/net#Error for
// more information on the net.Error interface.
func GetNetError(err error) (net.Error, bool) {
	var (
		netError net.Error
		ok       bool
	)

	for {
		if err == nil {
			break
		}

		netError, ok = err.(net.Error)
		if ok {
			break
		}

		err = errors.Unwrap(err)
	}

	return netError, ok
}
