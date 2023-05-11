package gadb

import (
	"testing"
)

func Test_transport_VerifyResponse(t *testing.T) {
	SetDebug(true)

	transport, err := NewTransport("localhost:5037")
	if err != nil {
		t.Fatal(err)
	}
	defer transport.Close()

	// err = Transport.Send("host:123version")
	err = transport.Send("host:version")
	if err != nil {
		t.Fatal(err)
	}

	err = transport.VerifyResponse()
	if err != nil {
		t.Fatal(err)
	}
}
