package tweet

import (
	"github.com/matryer/is"
	"testing"
)

func TestParseTweetURL(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		expectedErr  error
		name         string
		url          string
		expectedUser string
		expectedId   int64
	}{
		{
			name:         "a valid twitter url",
			url:          "https://twitter.com/Methodsco/status/1165711892163305478",
			expectedId:   1165711892163305478,
			expectedUser: "Methodsco",
			expectedErr:  nil,
		},
		{
			name:         "ends with ?s=20",
			url:          "https://twitter.com/headinthebox/status/1221307537804296192?s=20",
			expectedId:   1221307537804296192,
			expectedUser: "headinthebox",
			expectedErr:  nil,
		},
		{
			name:         "invalid url",
			url:          "https://rayyildiz.com",
			expectedId:   0,
			expectedUser: "",
			expectedErr:  ErrInvalidURL,
		},
		{
			name:         "invalid twitter url",
			url:          "https://twitter.com/headinthebox/INCORRECT_SUBPATH/1221307537804296192?s=20",
			expectedId:   0,
			expectedUser: "",
			expectedErr:  ErrInvalidURL,
		},
		{
			name:         "not start with http",
			url:          "twitter.com/rayyildiz/status/1218082900085813248",
			expectedId:   0,
			expectedUser: "",
			expectedErr:  ErrInvalidURL,
		},
		{
			name:         "support mobile subdomain",
			url:          "https://mobile.twitter.com/_denizparlak/status/1454140209306734602",
			expectedId:   1454140209306734602,
			expectedUser: "_denizparlak",
			expectedErr:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualId, actualUser, actualErr := parseTweetURL(test.url)
			is.Equal(test.expectedErr, actualErr)

			is.Equal(test.expectedId, actualId)
			is.Equal(test.expectedUser, actualUser)
		})
	}
}
