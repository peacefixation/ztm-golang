//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hours   int
	Minutes int
	Seconds int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(s string) (Time, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 3 {
		return Time{}, &TimeParseError{"invalid time format", s}
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid hours part: %v", parts[0]), s}
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid minutes part: %v", parts[1]), s}
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid seconds part: %v", parts[2]), s}
	}

	if hours < 0 || hours > 23 {
		return Time{}, &TimeParseError{"hours out of range 0 <= hours <= 23", fmt.Sprintf("%d", hours)}
	}

	if minutes < 0 || minutes > 59 {
		return Time{}, &TimeParseError{"minutes out of range 0 <= minutes <= 59", fmt.Sprintf("%d", minutes)}
	}

	if seconds < 0 || seconds > 59 {
		return Time{}, &TimeParseError{"seconds out of range 0 <= seconds <= 59", fmt.Sprintf("%d", seconds)}
	}

	return Time{
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}, nil
}
