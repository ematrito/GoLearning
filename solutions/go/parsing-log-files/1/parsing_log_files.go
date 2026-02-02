package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^(\[INF\]|\[ERR\])`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"[^"]*(?i)password[^"]*"`)
    totalCount := 0
    for _, line := range lines {

		matches := re.FindAllString(line, -1)
		totalCount += len(matches)
	}

	return totalCount
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(?i)end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}


func TagWithUserName(lines []string) []string {
re := regexp.MustCompile(`User\s+([^\s]+)`)
	
	result := make([]string, len(lines))

	for i, line := range lines {

		matches := re.FindStringSubmatch(line)

		if len(matches) > 1 {

			result[i] = "[USR] " + matches[1] + " " + line
		} else {
			result[i] = line
		}
	}

	return result
}
