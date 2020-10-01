package dsmr

import "errors"

var (
	RunUnknownObis        = errors.New("unknown obis code received")
	ErrIdentifierNotFound = errors.New("no Obis identifier found in file")
	ErrNoDeviceProvided   = errors.New("no device provided")
)
