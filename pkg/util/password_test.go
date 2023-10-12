package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hash, err := HashPassword(password)
	require.NoError(t, err)
	require.NotZero(t, len(hash))
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password123"
	hash, _ := HashPassword(password)

	var match bool
	var err error

	if match, err = CheckPasswordHash(password, hash); err != nil || !match {
		t.Errorf("CheckPasswordHash() returned false for a valid password: %s", err)
	}

	if match, err = CheckPasswordHash("wrong"+password, hash); err == nil && match {
		t.Errorf("CheckPasswordHash() returned true for a invvalid password: %s", err)
	}
}
