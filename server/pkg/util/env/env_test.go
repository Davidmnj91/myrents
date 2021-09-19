package env

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvAsStringOrFallback(t *testing.T) {
	const expected = "foo"

	assert := assert.New(t)

	key := "SET_TEST_VAR"
	os.Setenv(key, expected)
	assert.Equal(expected, GetEnvAsStringOrFallback(key, "~"+expected))

	key = "UNSET_TEST_VAR"
	assert.Equal(expected, GetEnvAsStringOrFallback(key, expected))
}

func TestGetEnvAsIntOrFallback(t *testing.T) {
	const expected = 1

	assert := assert.New(t)

	key := "SET_TEST_VAR"
	os.Setenv(key, strconv.Itoa(expected))
	returnVal, _ := GetEnvAsIntOrFallback(key, 1)
	assert.Equal(expected, returnVal)

	key = "UNSET_TEST_VAR"
	returnVal, _ = GetEnvAsIntOrFallback(key, expected)
	assert.Equal(expected, returnVal)

	key = "SET_TEST_VAR"
	os.Setenv(key, "not-an-int")
	returnVal, err := GetEnvAsIntOrFallback(key, 1)
	assert.Equal(expected, returnVal)
	if err == nil {
		t.Error("expected error")
	}
}
