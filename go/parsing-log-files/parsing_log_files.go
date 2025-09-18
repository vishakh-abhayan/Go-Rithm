package parsinglogfiles

import "regexp"

var (
	validLineRegex    = regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	splitLogLineRegex = regexp.MustCompile(`<[~*=-]*>`)
	passwordsRegex    = regexp.MustCompile(`".*(?i:password).*"`)
	endOfLineRegex    = regexp.MustCompile(`end-of-line\d+`)
	userNameRegex     = regexp.MustCompile(`User\s+(\S+)`)
)

func IsValidLine(text string) bool {
	return validLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogLineRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	for _, x := range lines {
		if passwordsRegex.MatchString(x) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		if matches := userNameRegex.FindStringSubmatch(line); matches != nil {
			username := matches[1]
			lines[i] = "[USR] " + username + " " + line
		}
	}
	return lines
}
