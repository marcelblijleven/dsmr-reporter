package dsmr_test

import (
	"errors"
	"github.com/marcelblijleven/dsmrreporter/dsmr"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
	"testing"
)

type TestPortReader struct {
	openErr       error
	OpenWasCalled bool
}

func (t *TestPortReader) OpenPort(_ *serial.Config) (*serial.Port, error) {
	t.OpenWasCalled = true

	if t.openErr != nil {
		return nil, t.openErr
	}

	return &serial.Port{}, nil
}

func TestOpenPort(t *testing.T) {
	portReader := &TestPortReader{OpenWasCalled: false}

	var (
		baud   = ""
		parity = ""
		device = "/tty/test"
	)

	_, err := dsmr.OpenPort(portReader.OpenPort, device, baud, parity)
	assert.NoError(t, err)
	assert.True(t, portReader.OpenWasCalled)
}

func TestOpenPort_EmptyDevice(t *testing.T) {
	portReader := &TestPortReader{OpenWasCalled: false}

	var (
		baud   = ""
		parity = ""
		device = ""
	)

	_, err := dsmr.OpenPort(portReader.OpenPort, device, baud, parity)
	assert.Error(t, err)
	assert.ErrorIs(t, dsmr.ErrNoDeviceProvided, err)
	assert.False(t, portReader.OpenWasCalled)
}

func TestOpenPort_PortReaderReturnsErr(t *testing.T) {
	expectedErr := errors.New("error reading tty")

	portReader := &TestPortReader{
		OpenWasCalled: false,
		openErr:       expectedErr,
	}

	var (
		baud   = ""
		parity = ""
		device = "/tty/test"
	)

	_, err := dsmr.OpenPort(portReader.OpenPort, device, baud, parity)
	assert.Error(t, err)
	assert.ErrorIs(t, expectedErr, err)
	assert.True(t, portReader.OpenWasCalled)
}
