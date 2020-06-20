package tweet

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
)

func parseTweetURL(url string) (int64, string, error) {
	r, err := regexp.Compile("(/twitter.com/)(.*)(/status/)([0-9]*)")
	if err != nil {
		return 0, "", err
	}

	parts := r.FindStringSubmatch(url)
	if len(parts) < 5 {
		return 0, "", ErrInvalidURL
	}

	idStr := parts[4]
	user := parts[2]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, "", err
	}

	return id, user, nil
}
