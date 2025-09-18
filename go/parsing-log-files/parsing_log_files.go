package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`".*[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd].*"`)
	for _, x := range lines {
		if re.MatchString(x) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	findUser := regexp.MustCompile(`User\s`)
	findUserId := regexp.MustCompile(`[a-zA-Z]+\d+`)

	for i, x := range lines {
		if findUser.MatchString(x) {
			state := lines[i]
			lines[i] = "[USR] " + findUserId.FindString(x) + " " + state
		}

	}
	return lines
}
