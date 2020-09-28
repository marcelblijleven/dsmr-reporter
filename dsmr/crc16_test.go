package dsmr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateCRC16Checksum(t *testing.T) {
	input := []byte("smartmeter")
	expected := "8BC5"
	actual := GenerateCRC16Checksum(input)
	assert.Equal(t, expected, actual)
}
