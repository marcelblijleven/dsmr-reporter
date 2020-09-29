package dsmr

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsHeader(t *testing.T) {
	assert.True(t, isHeader("/c0ffee"))
	assert.False(t, isHeader("c0ffee"))
}

func TestIsEmptyOrEnd(t *testing.T) {
	assert.True(t, isEmptyOrEnd(""))
	assert.True(t, isEmptyOrEnd("!"))
	assert.False(t, isEmptyOrEnd("c0ffee"))
}

func TestIsMultipartValue(t *testing.T) {
	_, ok := isMultipartValue("(5)(0-0:96.7.19)(170502102344S)(0000003590*s)(170502102349S)(0000003590*s)(170522124127S)(0000001987*s)(170728124311S)(0000016058*s)(191114125021W)(0000003300*s)")
	assert.True(t, ok)
	_, ok = isMultipartValue("(234.0*V)")
	assert.False(t, ok)
}

func TestGetObisCode(t *testing.T) {
	code, err := getObisCode("1-0:1.8.1(002137.886*kWh)")
	assert.Nil(t, err)
	assert.Equal(t, "1-0:1.8.1", code)

	_, err = getObisCode("no-obis-code")

}

func TestGetValueUnit(t *testing.T) {
	line := "(002137.886*kWh)"
	expectedValue := "002137.886"
	expectedUnit := "kWh"
	val, unit := getValueUnit(line)

	assert.Equal(t, expectedValue, val)
	assert.Equal(t, expectedUnit, unit)
}

func TestGetValueUnit_NoUnit(t *testing.T) {
	line := "(002137.886)"
	expectedValue := "002137.886"
	expectedUnit := ""
	val, unit := getValueUnit(line)

	assert.Equal(t, expectedValue, val)
	assert.Equal(t, expectedUnit, unit)
}

func TestCleanTimestamp(t *testing.T) {
	s := "200505212756S"
	exp := "200505212756"
	assert.Equal(t, exp, cleanTimestamp(s))
}

func TestParseTimestamp(t *testing.T) {
	d, err := parseTimestamp("200105212756S", "")

	assert.Nil(t, err)
	assert.Equal(t, "Europe/Amsterdam", d.Location().String())
	assert.Equal(t, 2020, d.Year())
	assert.Equal(t, time.Month(1), d.Month())
	assert.Equal(t, 5, d.Day())
}

func TestParseTimestamp_IncorrectLocation(t *testing.T) {
	_, err := parseTimestamp("200105212756S", "Westeros/KingsLanding")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unknown time zone")
}

func TestParseTimestamp_IncorrectTimestamp(t *testing.T) {
	_, err := parseTimestamp("133319993", "Europe/Amsterdam")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "month out of range")
}
