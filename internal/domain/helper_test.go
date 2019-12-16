package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunc(t *testing.T) {
	require.NotEmpty(t, GetFunc())
}
