// Package bob responds to different remarks
package bob

import "strings"

// Hey takes a string (remark) and replies based on certain characteristics of that string
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == strings.ToUpper(remark) && strings.ToUpper(remark) != strings.ToLower(remark) {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
