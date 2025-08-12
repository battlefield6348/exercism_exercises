package parsinglogfiles

import (
    "regexp"
    "strings"
    "fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[[TRC|DBG|INF|WRN|ERR|FTL]+\]\w*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<\W*\>`)
	sl := re.Split(text, -1)
	return sl
}

func CountQuotedPasswords(lines []string) int {
	// 匹配包含單字 "password" 的雙引號字串（不分大小寫）
	re := regexp.MustCompile(`(?i)"[^"]*\bpassword\b[^"]*"`)
	count := 0
	for i := range lines {
		if re.MatchString(lines[i]) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+[A-Za-z0-9]+`)
	reLines := []string{}
	for i := range lines {
		userString := re.FindString(lines[i])
		if userString == "" {
			reLines = append(reLines, lines[i])
		} else {
			splitUserString := strings.Split(userString, " ")
			reLines = append(reLines, fmt.Sprintf("[USR] %s %s", splitUserString[len(splitUserString)-1], lines[i]))
		}
	}
	return reLines
}
