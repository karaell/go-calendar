package events

import (
	"regexp"
)

const pattern = "^[a-zA-Z0-9 _,-\\.]{3,50}$"

func isValidTitle(title string) bool {
	matched, _ := regexp.MatchString(pattern, title)

	return matched
}
