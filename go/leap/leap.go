package leap

// IsLeapYear returns true if the given year was a leap year
func IsLeapYear(year int) bool {

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false

}
