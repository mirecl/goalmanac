package validate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCreateValidater - создание инстанции Валидаторов
func TestCreateValidater(t *testing.T) {
	err := createValidate()
	require.NoError(t, err)
}
