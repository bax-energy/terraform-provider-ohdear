package ohdear

import (
	"fmt"
	"time"
)

// Time wraps time.Time to handle Oh Dear's "YYYY-MM-DD HH:MM:SS" format
type Time struct {
	time.Time
}

// UnmarshalJSON implements custom unmarshaling for Oh Dear timestamps
func (t *Time) UnmarshalJSON(b []byte) error {
	s := string(b)
	// Remove quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" || s == "null" {
		t.Time = time.Time{}
		return nil
	}

	// Parse using Oh Dear format
	const layout = "2006-01-02 15:04:05"
	parsed, err := time.Parse(layout, s)
	if err != nil {
		return fmt.Errorf("failed to parse Oh Dear time %q: %w", s, err)
	}
	t.Time = parsed
	return nil
}

// MarshalJSON to keep the same format when sending back
func (t Time) MarshalJSON() ([]byte, error) {
	const layout = "2006-01-02 15:04:05"
	return []byte(`"` + t.Format(layout) + `"`), nil
}
