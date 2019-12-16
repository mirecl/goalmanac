package validate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCreateValidater - создание инстанции Валидаторов + валидация
func TestValidaterSuccess(t *testing.T) {
	err := createValidate()
	require.NoError(t, err)

	result, err := Create.Validate([]byte(`{"title":"user","user":"user","body":"user","start":"2019-12-20T11:40:00.000Z","duration":"30m"}`))

	require.NoError(t, err)
	require.Equal(t, 0, len(result.Errors()))
}

func TestValidaterBad(t *testing.T) {
	err := createValidate()
	require.NoError(t, err)

	result, err := Create.Validate([]byte(`{"title":"user","user":"user","body":"user","start":"2019-12-20T11:40:00.000Z","duration":"11m"}`))

	require.NoError(t, err)
	require.Equal(t, 1, len(result.Errors()))
}
