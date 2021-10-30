package tweet

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
)

// re represents a compiled regular expression for matching twitter URLs.
var re = regexp.MustCompile(`((/twitter.com/)|(/mobile.twitter.com/))(.*)(/status/)([0-9]*)`)

// parseTweetURL parses a tweet URL into its components ( userId, tweetId and error).
func parseTweetURL(url string) (int64, string, error) {
	parts := re.FindStringSubmatch(url)
	if len(parts) < 7 {
		return 0, "", ErrInvalidURL
	}

	idStr := parts[6]
	user := parts[4]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, "", err
	}

	return id, user, nil
}
