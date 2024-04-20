package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// Returns true if line starts with
//

// var validTag []string = []string{
// 	"[TRC]",
// 	"[DBG]",
// 	"[INF]",
// 	"[WRN]",
// 	"[ERR]",
// 	"[FTL]",
// }

// IsValidLine returns true if string starts with validTag
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	//fmt.Println("Match found:", re.FindString(text))
	return re.MatchString(text)
}

// SplitLogLine returns string split by token starting with '<' and ending with
// '>' with optionally any of the following chars in between: '~*=-'
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]+>`)
	//fmt.Println("Match found:", re.FindAllString(text, -1))
	return re.Split(text, -1)
}

// Count any lines that contain "password", somewhere between quotes
// and caseinsensitive
func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*(password).*"`)
	count := 0
	for _, line := range lines {
		match := re.MatchString(line)
		if match {
			// fmt.Println("Match:", re.FindString(line))
			count++
		}
	}
	return count
}

// RemoveEndOfLineText removes string "end-of-line" followed by any numbers fro the
// input string
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d*)`)
	str := re.ReplaceAllString(text, "")
	//fmt.Println("Replaced text:", str)
	return str
}

// TagWithUserName tags a line that contains a "User" followed by
// any number of spaces and a username with the tag: "[USR] <username>"
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User(\s+)([[:alnum:]]+)`)
	for idx, line := range lines {
		if re.MatchString(line) {
			match := re.FindStringSubmatch(line)
			//fmt.Println("match:", match[2])
			name := match[2]
			lines[idx] = fmt.Sprintf("[USR] %s %s", name, lines[idx])
		}
	}
	return lines
}
