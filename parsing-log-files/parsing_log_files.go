package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	// panic("Please implement the IsValidLine function")
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	matched := re.MatchString(text)
	return matched
}

func SplitLogLine(text string) []string {
	// panic("Please implement the SplitLogLine function")
	re := regexp.MustCompile(`<[\-~*=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	// panic("Please implement the CountQuotedPasswords function")
	count := 0
	re := regexp.MustCompile(`"(?i).*password.*"`)

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	// panic("Please implement the RemoveEndOfLineText function")
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// panic("Please implement the TagWithUserName function")
	re := regexp.MustCompile(`User\s+(\S+)`)
	var result []string

	for _, line := range lines {
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)
			userName := matches[1]
			line = "[USR] " + userName + " " + line
		}
		result = append(result, line)
	}
	return result
}
