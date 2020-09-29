package dsmr

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("./testdata/example_landis_gyr_E350.txt")

	if err != nil {
		t.Fatal("error opening test data file")
	}

	br := bufio.NewReader(f)
	reporter := &TestReporter{
		UpdateWasCalled: false,
		Logs:            make([]string, 0),
		Errors:          make([]error, 0),
	}

	Read(br, reporter, "")
	assert.True(t, reporter.UpdateWasCalled)
	assert.Len(t, reporter.Errors, 0)
}

func TestRead_PartialMessage(t *testing.T) {
	f, err := os.Open("./testdata/example_landis_gyr_E350_partial_message.txt")

	if err != nil {
		t.Fatal("error opening test data file")
	}

	br := bufio.NewReader(f)
	reporter := &TestReporter{
		UpdateWasCalled: false,
		Logs:            make([]string, 0),
		Errors:          make([]error, 0),
	}

	Read(br, reporter, "")
	fmt.Println(reporter)
	assert.False(t, reporter.UpdateWasCalled)
	assert.Equal(t, 1, len(reporter.Logs))
	assert.Equal(t, "EOF reached", reporter.Logs[0])
}

func TestRead_PartialMessage_PeakAtEnd(t *testing.T) {
	buf := bytes.NewBufferString("")
	buf.Truncate(0)
	br := bufio.NewReader(buf)
	reporter := &TestReporter{
		UpdateWasCalled: false,
		Logs:            make([]string, 0),
		Errors:          make([]error, 0),
	}

	Read(br, reporter, "")
	assert.False(t, reporter.UpdateWasCalled)
	assert.Equal(t, 1, len(reporter.Logs))
}

func TestRead_IncorrectChecksum(t *testing.T) {
	f, err := os.Open("./testdata/example_landis_gyr_E350_incorrect_checksum.txt")

	if err != nil {
		t.Fatal("error opening test data file")
	}

	br := bufio.NewReader(f)
	reporter := &TestReporter{
		UpdateWasCalled: false,
		Logs:            make([]string, 0),
		Errors:          make([]error, 0),
	}

	Read(br, reporter, "")

	assert.False(t, reporter.UpdateWasCalled)
	assert.True(t, len(reporter.Errors) == 1)
	assert.Contains(t, reporter.Errors[0].Error(), "crc does not match")
}

func TestRead_ErrorParsingLine(t *testing.T) {
	f, err := os.Open("./testdata/example_landis_gyr_E350_with_incorrect_obis.txt")

	if err != nil {
		t.Fatal("error opening test data file")
	}

	br := bufio.NewReader(f)
	reporter := &TestReporter{
		UpdateWasCalled: false,
		Logs:            make([]string, 0),
		Errors:          make([]error, 0),
	}
	Read(br, reporter, "")
	assert.False(t, reporter.UpdateWasCalled)
	assert.True(t, len(reporter.Errors) == 1)
	assert.Contains(t, reporter.Errors[0].Error(), "unknown obis code received")
}

type TestReporter struct {
	UpdateWasCalled bool
	Logs            []string
	Errors          []error
}

func (t *TestReporter) Update(_ Telegram) {
	t.UpdateWasCalled = true
}

func (t *TestReporter) Log(msg string) {
	t.Logs = append(t.Logs, msg)
}

func (t *TestReporter) Error(err error) {
	t.Errors = append(t.Errors, err)
}
