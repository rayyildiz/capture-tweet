package convert

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	t1 := "test"
	t2 := "hello"

	require.Equal(t, &t1, String(t1))
	require.Equal(t, &t2, String(t2))
	require.Nil(t, String(""))
}

func TestInt(t *testing.T) {
	t1 := 1
	t2 := 2

	require.Equal(t, &t1, Int(t1))
	require.Equal(t, &t2, Int(t2))
}

func TestTime(t *testing.T) {
	t1 := time.Now()
	t2 := time.Now().Add(20 * time.Hour)

	require.Equal(t, &t1, Time(t1))
	require.Equal(t, &t2, Time(t2))
}
