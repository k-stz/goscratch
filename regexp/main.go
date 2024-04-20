package main

import (
	"fmt"

	. "github.com/k-stz/goscratch/regexp/parsinglogfiles"
)

func main() {
	fmt.Println(IsValidLine("[ERR] A good error here"))
	// => true
	fmt.Println(IsValidLine("Any old [ERR] text"))
	// => false
	fmt.Println(IsValidLine("[BOB] Any old text"))
	// => false
	sll := SplitLogLine("section 1<*>section 2<~~~>section 3<**>section 4")
	// => []string{"section 1", "section 2", "section 3"},
	fmt.Println(sll)
	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	fmt.Println("CQP:", CountQuotedPasswords(lines))
	RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27")
	// => "[INF]  Network Failure "
	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})
	// => []string {
	//  "[USR] James123 [WRN] User James123 has exceeded storage space.",
	//  "[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
	//  "[INF] Users can login again after 23:00.",
	//  "[DBG] We need to check that user names are at least 6 chars long."
	// }
	fmt.Println("Result TagWithUserName:")
	for _, line := range result {
		fmt.Println(line)
	}
}
