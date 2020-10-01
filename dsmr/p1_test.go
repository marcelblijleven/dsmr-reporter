package dsmr

import (
	"errors"
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

	_, err := OpenPort(portReader.OpenPort, device, baud, parity)
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

	_, err := OpenPort(portReader.OpenPort, device, baud, parity)
	assert.Error(t, err)
	assert.ErrorIs(t, ErrNoDeviceProvided, err)
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

	_, err := OpenPort(portReader.OpenPort, device, baud, parity)
	assert.Error(t, err)
	assert.ErrorIs(t, expectedErr, err)
	assert.True(t, portReader.OpenWasCalled)
}

func TestCheckParity(t *testing.T) {
	// even mark none odd space
	even, err := checkParity("even")
	assert.NoError(t, err)
	assert.Equal(t, serial.ParityEven, even)

	mark, err := checkParity("mark")
	assert.NoError(t, err)
	assert.Equal(t, serial.ParityMark, mark)

	none, err := checkParity("none")
	assert.NoError(t, err)
	assert.Equal(t, serial.ParityNone, none)

	odd, err := checkParity("odd")
	assert.NoError(t, err)
	assert.Equal(t, serial.ParityOdd, odd)

	space, err := checkParity("space")
	assert.NoError(t, err)
	assert.Equal(t, serial.ParitySpace, space)

	p, err := checkParity("")
	assert.NoError(t, err)
	assert.Equal(t, DefaultParity, p)

	_, err = checkParity("uneven")
	assert.Error(t, err)
	assert.EqualError(t, err, "could not parse \"uneven\" into parity")
}

func TestCheckBaud(t *testing.T) {
	b, err := checkBaud("")
	assert.NoError(t, err)
	assert.Equal(t, DefaultBaud, b)

	b, err = checkBaud("1337")
	assert.NoError(t, err)
	assert.Equal(t, 1337, b)

	_, err = checkBaud("baud?")
	assert.Error(t, err)
	assert.EqualError(t, err, "strconv.Atoi: parsing \"baud?\": invalid syntax")
}

func TestGetP1Config(t *testing.T) {
	// No device provided
	_, err := getP1Config("", "", "")
	assert.Error(t, err)
	assert.ErrorIs(t, ErrNoDeviceProvided, err)

	// Invalid baud provided
	_, err = getP1Config("/dev/ttyUSB0", "baud?", "")
	assert.Error(t, err)
	assert.EqualError(t, err, "strconv.Atoi: parsing \"baud?\": invalid syntax")

	// Invalid parity provided
	_, err = getP1Config("/dev/ttyUSB0", "", "uneven")
	assert.Error(t, err)
	assert.EqualError(t, err, "could not parse \"uneven\" into parity")

	config, err := getP1Config("my-device", "1337", "odd")

	if err == nil {
		assert.Equal(t, "my-device", config.Name)
		assert.Equal(t, serial.ParityOdd, config.Parity)
	} else {
		assert.Fail(t, "expected err to be non nil")
	}
}
