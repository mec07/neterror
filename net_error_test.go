package neterror_test

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mec07/neterror"
	"gotest.tools/assert"
)

func TestGetNetError(t *testing.T) {
	table := []struct {
		name          string
		err           error
		shouldSucceed bool
	}{
		{
			name:          "top level DNS error",
			err:           &net.DNSError{},
			shouldSucceed: true,
		},
		{
			name:          "wrapped DNS error",
			err:           fmt.Errorf("stuff to wrap with: %w", &net.DNSError{}),
			shouldSucceed: true,
		},
		{
			name:          "not a network error",
			err:           errors.New("hello world"),
			shouldSucceed: false,
		},
		{
			name:          "nil error",
			err:           nil,
			shouldSucceed: false,
		},
	}

	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			_, ok := neterror.GetNetError(test.err)
			assert.Equal(t, test.shouldSucceed, ok)
		})
	}
}

func ExampleGetNetError() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		w.WriteHeader(200)
	})
	server := httptest.NewServer(handler)

	client := http.Client{Timeout: time.Millisecond}
	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		fmt.Printf("Unexpected error: %v", err)
		return
	}

	_, err = client.Do(req)
	netError, ok := neterror.GetNetError(err)
	if !ok {
		fmt.Println("Expected a net.Error")
	}

	if netError.Temporary() && netError.Timeout() {
		fmt.Println("Temporary Timeout error")
	}
	// Output: Temporary Timeout error
}
