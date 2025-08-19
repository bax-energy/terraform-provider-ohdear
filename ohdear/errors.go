// errors.go
package ohdear

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	ErrValidation = errors.New("validation error")
	ErrHTTPStatus = errors.New("http status error")
)

type ValidationError struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func (e *ValidationError) Error() string {
	if e == nil {
		return ""
	}
	if len(e.Errors) == 0 {
		return e.Message
	}
	return fmt.Sprintf("%s (%s)", e.Message, flattenFieldErrors(e.Errors))
}

// HTTPStatusError captures non-validation HTTP failures (status + snippet).
type HTTPStatusError struct {
	StatusCode int
	Status     string
	Body       string // trimmed snippet of the response body (may be empty)
}

func (e *HTTPStatusError) Error() string {
	if e == nil {
		return ""
	}
	if e.Body != "" {
		return fmt.Sprintf("%d %s: %s", e.StatusCode, e.Status, e.Body)
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.Status)
}

// parseAPIError reads a non-2xx response and returns a wrapped static error
func parseAPIError(resp *http.Response) error {
	defer resp.Body.Close()

	// Try validation error shape first.
	var v ValidationError
	if err := json.NewDecoder(resp.Body).Decode(&v); err == nil && v.Message != "" {
		// Wrap the typed error with the sentinel (err113-friendly).
		return fmt.Errorf("%w: %w", ErrValidation, &v)
	}

	// Not a validation payload — produce an HTTPStatusError with a short body snippet.
	snippetBytes, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
	snippet := strings.TrimSpace(string(snippetBytes))

	return fmt.Errorf("%w: %w", ErrHTTPStatus, &HTTPStatusError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       snippet,
	})
}

// flattenFieldErrors compresses map[field][]messages to "field: m1; m2, field2: m1".
func flattenFieldErrors(m map[string][]string) string {
	if len(m) == 0 {
		return ""
	}
	parts := make([]string, 0, len(m))
	for field, msgs := range m {
		if len(msgs) == 0 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", field, strings.Join(msgs, "; ")))
	}
	return strings.Join(parts, ", ")
}
