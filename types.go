package errutil

import (
	"errors"
	"fmt"
	"time"
)

var ErrRateLimited = errors.New("rate limited")

// ValidationError represents a validation error with field context
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(field string, value interface{}, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

// ConfigError represents configuration-related errors
type ConfigError struct {
	Operation string
	Path      string
	Cause     error
}

func (e ConfigError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("config %s failed for %s: %v", e.Operation, e.Path, e.Cause)
	}
	return fmt.Sprintf("config %s failed for %s", e.Operation, e.Path)
}

func (e ConfigError) Unwrap() error {
	return e.Cause
}

// NewConfigError creates a new configuration error
func NewConfigError(operation, path string, cause error) ConfigError {
	return ConfigError{
		Operation: operation,
		Path:      path,
		Cause:     cause,
	}
}

// DiscordError represents Discord API related errors
type DiscordError struct {
	Operation string
	Code      int
	Message   string
	Cause     error
}

func (e DiscordError) Error() string {
	if e.Code > 0 {
		return fmt.Sprintf("Discord API error during %s (code %d): %s", e.Operation, e.Code, e.Message)
	}
	return fmt.Sprintf("Discord API error during %s: %s", e.Operation, e.Message)
}

func (e DiscordError) Unwrap() error {
	return e.Cause
}

// NewDiscordError creates a new Discord API error
func NewDiscordError(operation string, code int, message string, cause error) DiscordError {
	return DiscordError{
		Operation: operation,
		Code:      code,
		Message:   message,
		Cause:     cause,
	}
}

// IsRetryableError determines if an error can be retried
func IsRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Check for specific retryable errors
	if errors.Is(err, ErrRateLimited) {
		return true
	}

	// Check for Discord errors that might be retryable
	var discordErr DiscordError
	if errors.As(err, &discordErr) {
		// 5xx errors are typically retryable
		return discordErr.Code >= 500 && discordErr.Code < 600
	}

	return false
}

type RetryManager struct {
	maxRetries    int
	baseDelay     time.Duration
	maxDelay      time.Duration
	backoffFactor float64
}

func NewRetryManager() *RetryManager {
	return &RetryManager{
		maxRetries:    3,
		baseDelay:     time.Second * 2,
		maxDelay:      time.Minute,
		backoffFactor: 2.0,
	}
}
