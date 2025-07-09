package errors

// ErrorMapping represents a description and fix for a known error.
type ErrorMapping struct {
	Description string `json:"description"`
	Fix         string `json:"fix"`
}
