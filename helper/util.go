package helper

import (
	"errors"
	"regexp"
)

func MatchSingle(re *regexp.Regexp, content string) (string, error) {
	matched := re.FindAllStringSubmatch(content, -1)
	if len(matched) < 1 {
		return "", errors.New("errorNoMatched")
	}
	return matched[0][1], nil
}
