// Package validate provides simple syntactic validators for common
// user-input formats.
//
// Homework — Task 2 (Lesson 6: File I/O, JSON and Testing):
// Implement ValidateEmail and/or ValidatePhone below (your mentor may
// ask for just one) and extend the test tables in validate_test.go to
// at least 8 cases each, including edge cases.
package validate

import (
	"strings"
	"unicode"
)

// ValidateEmail reports whether s is a syntactically valid email address.
//
// TODO: implement this function. At minimum it should:
//   - reject the empty string,
//   - require exactly one "@" with a non-empty local part and domain part,
//   - reject values containing whitespace.
//
// Document any additional decisions you make (e.g. how you handle a
// trailing dot, consecutive dots, or unicode characters) in a comment
// here, and add matching test cases in validate_test.go.
func ValidateEmail(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if unicode.IsSpace(r) {
			return false
		}
	}
	at := strings.Index(s, "@")
	if at < 1 {
		return false
	}
	domain := s[at+1:]
	if domain == "" || strings.Count(s, "@") != 1 {
		return false
	}
	if !strings.Contains(domain, ".") {
		return false
	}
	return true
}

// ValidatePhone reports whether s is a syntactically valid phone number.
//
// TODO: implement this function. Decide which format(s) you accept
// (e.g. "+380501234567", "050-123-4567") and document your decision
// here. At minimum it should reject the empty string and any value
// containing letters.
func ValidatePhone(s string) bool {
	if s == "" {
		return false
	}
	digits := 0
	for i, r := range s {
		if r == '+' && i == 0 {
			continue
		}
		if r == '-' || r == ' ' || r == '(' || r == ')' {
			continue
		}
		if r >= '0' && r <= '9' {
			digits++
			continue
		}
		return false
	}
	return digits >= 7 && digits <= 15
}
