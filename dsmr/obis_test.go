package dsmr

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetObis(t *testing.T) {
	obis, err := GetObis("1-3:0.2.8")
	assert.Equal(t, "Version", obis.Name)
	assert.Nil(t, err)

	_, err = GetObis("does-not:exist")
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, RunUnknownObis))
}
