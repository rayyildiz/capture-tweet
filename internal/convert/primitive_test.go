package convert

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	is := is.New(t)

	t1 := "test"
	t2 := "hello"

	is.Equal(&t1, String(t1))
	is.Equal(&t2, String(t2))
	is.True(nil == String("")) // empty string is null
}

func TestInt(t *testing.T) {
	is := is.New(t)

	t1 := 1
	t2 := 2

	is.Equal(&t1, Int(t1))
	is.Equal(&t2, Int(t2))
}

func TestTime(t *testing.T) {
	is := is.New(t)

	t1 := time.Now()
	t2 := time.Now().Add(20 * time.Hour)

	is.Equal(&t1, Time(t1))
	is.Equal(&t2, Time(t2))
}

func TestInt64(t *testing.T) {
	is := is.New(t)

	var actual int64 = 5
	is.Equal(&actual, Int64(actual))
}
