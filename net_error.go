package neterror

import (
	"errors"
	"net"
)

// GetNetError recursively checks the chain of errors until it finds an error
// from the net package that fulfills the net.Error interface. If we find one
// that means its safe to say that a network error has occurred. See:
// https://godoc.org/net#Error for more information on the net.Error interface.
func GetNetError(err error) (net.Error, bool) {
	for {
		if err == nil {
			break
		}

		// As there are errors that are not from the net package
		// that satisfy the net.Error interface, we need to explicitly
		// ensure that the error does indeed come from the net package.
		switch netError := err.(type) {
		case *net.DNSConfigError, *net.DNSError, *net.AddrError, *net.InvalidAddrError, *net.OpError, *net.UnknownNetworkError:
			return netError.(net.Error), true
		}

		err = errors.Unwrap(err)
	}

	return nil, false
}
