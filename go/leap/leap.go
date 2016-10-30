// Package leap determines if a given year is a leap year.
package leap

const testVersion = 2

// IsLeapYear takes in a year and returns T/F if a leap year
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
