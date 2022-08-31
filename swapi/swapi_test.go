package swapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetApparances(t *testing.T) {
	arg := "Alderaan"
	arg2 := "plutao"

	qtde := GetApparances(arg)
	qtde2 := GetApparances(arg2)

	require.NotEmpty(t, qtde)
	require.Equal(t, 2, qtde)
	require.Equal(t, 0, qtde2)

}
