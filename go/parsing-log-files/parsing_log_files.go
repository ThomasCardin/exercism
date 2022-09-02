package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile("\\^[TRC]|^\\[DBG]|^\\[INF]|^\\[WRN]|^\\[ERR]|^\\[FTL]")
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile("\\<[\\*~=-]*>")
	return re.Split(text, 9)
}

func CountQuotedPasswords(lines []string) int {
	counter := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for i := 0; i < len(lines); i++ {
		if re.MatchString(lines[i]) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	ret := []string{}
	for _, line := range lines {
		found := re.FindStringSubmatch(line)
		if found != nil {
			line = fmt.Sprintf("[USR] %s %s", found[1], line)
		}
		ret = append(ret, line)
	}
	return ret
}
