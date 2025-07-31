// Package bob emulates a lackadaisical teenager.
package bob

import "strings"

// Hey returns various teenager-like responses to questions.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if strings.HasSuffix(remark, "?") {
		if strings.ToUpper(remark) == remark && remark != strings.ToLower(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if strings.ToUpper(remark) == remark && remark != strings.ToLower(remark) {
		return "Whoa, chill out!"
	}
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
