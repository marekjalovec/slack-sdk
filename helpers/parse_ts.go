package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParseTimestamp(ts string) (*time.Time, error) {
	r := regexp.MustCompile(`^([0-9]+)\.([0-9]+)$`)
	if !r.MatchString(ts) {
		return nil, fmt.Errorf("invalid timestamp format")
	}

	m := r.FindStringSubmatch(ts)
	sec, err := strconv.ParseInt(m[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp format: %w", err)
	}
	nsec, err := strconv.ParseInt(m[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp format: %w", err)
	}

	tm := time.Unix(sec, nsec)
	return &tm, nil
}
