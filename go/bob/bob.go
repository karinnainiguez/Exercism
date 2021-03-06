// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var response string

	if remark == "" {
		response = "Fine. Be that way!"
	} else if remark == strings.ToUpper(remark) && regexp.MatchString('\w*[a-zA-Z]\w*', remark) &&remark[len(remark)-1] == byte('?') {
		response = "Calm down, I know what I'm doing!"
	} else if remark == strings.ToUpper(remark) {
		response = "Whoa, chill out!"
	} else if remark[len(remark)-1] == byte('?') {
		response = "Sure."
	} else {
		response = "Whatever."
	}

	return response
}
