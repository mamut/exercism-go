package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	expr := regexp.MustCompile(`\A\[(TRC|DBG|INF|WRN|ERR|FTL)]`)

	return expr.MatchString(text)
}

func SplitLogLine(text string) []string {
	expr := regexp.MustCompile(`<[~*=-]*>`)

	return expr.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	expr := regexp.MustCompile(`(?i)"[\w\s]*password[\w\s]*"`)

	count := 0

	for _, line := range lines {
		if expr.MatchString(line) {
			count += 1
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	expr := regexp.MustCompile(`end-of-line\d+`)

	return expr.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	expr := regexp.MustCompile(`User\s+(\w+)`)

	for i, line := range lines {
		match := expr.FindStringSubmatch(line)

		if match != nil {
			lines[i] = fmt.Sprintf("[USR] %s %s", match[1], line)
		}
	}

	return lines
}
