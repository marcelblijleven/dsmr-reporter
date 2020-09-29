package dsmr

import (
	"regexp"
	"strings"
	"time"
	"unicode"
)

const (
	DefaultTimeZone = "Europe/Amsterdam"
	timestampFormat = "060102150405"
)

var (
	regexObis      = regexp.MustCompile(`(\d-\d:\d{1,2}.\d{1,2}.\d{1,2})`)
	regexMultipart = regexp.MustCompile(`(\(\d{12}[S|W]\)\(\d+\*s\))`)
)

// isHeader checks if the provided line starts with "/"
func isHeader(line string) bool {
	return strings.HasPrefix(line, "/")
}

// isEmptyOrEnd checks if the line is empty, or has prefix "!"
func isEmptyOrEnd(line string) bool {
	return len(line) == 0 || strings.HasPrefix(line, "!")
}

// isMultipartValue checks if the provided string has multiple sequences of timestamp and duration values
func isMultipartValue(s string) ([]string, bool) {
	m := regexMultipart.FindAllString(s, -1)

	return m, m != nil && len(m) > 1
}

// getObisCode returns an OBIS identifier if found in the provided string
func getObisCode(s string) (string, error) {
	m := regexObis.FindStringSubmatch(s)

	if m == nil {
		return "", IdentifierNotFoundErr
	}

	return m[1], nil
}

// getValueUnit returns the value and optional unit from the provided value string
func getValueUnit(s string) (string, string) {
	val := strings.Split(strings.SplitAfter(s, "(")[1], ")")[0]
	parts := strings.Split(val, "*")

	// Check if the value has the optional unit
	if len(parts) == 2 {
		return parts[0], parts[1]
	}

	return val, ""
}

// cleanTimestamp trims any non numeric characters at the end of the string and returns the remaining string
func cleanTimestamp(ts string) string {
	cleaned := strings.TrimRightFunc(ts, func(r rune) bool {
		return !unicode.IsNumber(r)
	})

	return cleaned
}

// parseTimestamp parses the timestamp string in YYMMDDHHmmss format to time.Time
func parseTimestamp(ts, tz string) (time.Time, error) {
	if tz == "" {
		tz = DefaultTimeZone
	}

	t := time.Time{}
	ts = cleanTimestamp(ts)
	loc, err := time.LoadLocation(tz)

	if err != nil {
		return t, err
	}

	t, err = time.ParseInLocation(timestampFormat, ts, loc)

	if err != nil {
		return t, err
	}

	return t, nil
}
