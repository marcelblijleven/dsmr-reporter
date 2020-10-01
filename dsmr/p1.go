package dsmr

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"strconv"
)

const (
	DefaultBaud   = 115200
	DefaultParity = serial.ParityNone
)

// OpenPort opens a port on the provided serial device using the provided settings and returns a *bufio.Reader
func OpenPort(portReader func(c *serial.Config) (*serial.Port, error), device, baud, parity string) (*bufio.Reader, error) {
	config, err := getP1Config(device, baud, parity)

	if err != nil {
		return nil, err
	}

	p, err := portReader(config)

	if err != nil {
		return nil, err
	}

	return bufio.NewReader(p), nil
}

// checkBaud converts the provided baud string into an int. Defaults to 115200
func checkBaud(baud string) (int, error) {
	if baud == "" {
		return DefaultBaud, nil
	}

	b, err := strconv.Atoi(baud)

	if err != nil {
		return 0, err
	}

	return b, nil
}

// checkParity converts the provided parity string into a serial.Parity byte
func checkParity(parity string) (serial.Parity, error) {
	if parity == "" {
		return DefaultParity, nil
	}
	// even mark none odd space
	switch parity {
	case "even":
		return serial.ParityEven, nil
	case "mark":
		return serial.ParityMark, nil
	case "none":
		return serial.ParityNone, nil
	case "odd":
		return serial.ParityOdd, nil
	case "space":
		return serial.ParitySpace, nil
	default:
		return DefaultParity, fmt.Errorf("could not parse %q into parity", parity)
	}
}

// getP1Config converts the provided device, baud and parity strings into serial.Config
func getP1Config(device, baud, parity string) (*serial.Config, error) {
	if device == "" {
		return nil, ErrNoDeviceProvided
	}

	b, err := checkBaud(baud)

	if err != nil {
		return nil, err
	}

	p, err := checkParity(parity)

	if err != nil {
		return nil, err
	}

	return &serial.Config{
		Name:   device,
		Baud:   b,
		Parity: p,
	}, nil
}
