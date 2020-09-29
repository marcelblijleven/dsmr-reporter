package dsmr

import "errors"

var (
	UnknownObisCodeErr    = errors.New("unknown obis code received")
	IdentifierNotFoundErr = errors.New("no Obis identifier found in file")
)
