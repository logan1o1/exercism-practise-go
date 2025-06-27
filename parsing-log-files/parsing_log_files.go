package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)\"+[^"]*\bpassword\b\"+`)
	var total int
	for _, l := range lines {
		if re.MatchString(l) {
			total++
		}
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line[\d]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`(User[\s]+){1}([^\s]\w+)`)
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", match[2], line)
		}
	}
	return lines
}
