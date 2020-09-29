package dsmr

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestParseTelegram(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Amsterdam")
	f, err := os.Open("./testdata/example_landis_gyr_E350.txt")

	if err != nil {
		t.Fatal("error opening test data file")
	}

	msg, err := bufio.NewReader(f).ReadBytes('!')

	if err != nil {
		t.Fatal("error reading test data file")
	}

	telegram, err := ParseTelegram(msg, "")

	assert.NoErrorf(t, err, "expected err to be nil, got %q", err)
	assert.Equal(
		t, "50",
		telegram.Version,
		"expected telegram version to equal 50, got %q", telegram.Version,
	)

	expectedEquipmentID := "1530303136303031373837351232313131"
	assert.Equal(
		t,
		expectedEquipmentID,
		telegram.EquipmentID,
		"expected telegram equipment id to equal %q, got %q", expectedEquipmentID, telegram.EquipmentID,
	)

	controlTime := time.Date(2020, time.May, 5, 21, 27, 56, 0, loc)
	assert.Equal(t, controlTime, telegram.DateTime)
}

func TestParseTelegram_NoObisOnLine(t *testing.T) {
	data := []byte("no_obis_here\n")
	_, err := ParseTelegram(data, "")

	assert.Error(t, err)
}

func TestParseTelegram_UnknownObisOnLine(t *testing.T) {
	data := []byte("0-0:13.37.0(0002)\n")
	_, err := ParseTelegram(data, "")

	assert.Error(t, err)
}

func TestParseTelegram_IncorrectTimestampOnLine(t *testing.T) {
	data := []byte("0-0:1.0.0(201305212756S)\n")
	_, err := ParseTelegram(data, "")

	assert.Error(t, err)
}
