package errutil

const (
	// Error messages
	ErrValidationFailed           = "validation failed"
	ErrConfigOperationFailed      = "configuration operation failed"
	ErrDiscordOperationFailed     = "discord operation failed"
	ErrNonRetryable               = "non-retryable error encountered"
	ErrGlobalLoggerNotInitialized = "global logger not initialized for error handler"
	ErrOnAttempt                  = "error on attempt %d for %s"
	ErrOperationAttemptsFailed    = "operation %s failed after %d attempts. Last error: %w"

	// Constants for log messages
	MsgOperationRetrying            = "operation failed, retrying"
	MsgOperationSucceededAfterRetry = "operation succeeded after retry"
	MsgOperationFailedAllRetries    = "operation failed after all retries"
	MsgOperationFailedCleanup       = "operation failed, running cleanup"
	ErrOperationFailed              = "operation failed"
	MsgCriticalOperationFailed      = "critical operation failed"

	// Constants for error format strings
	ErrFmtNonRetryable                 = "non-retryable error in %s: %w"
	ErrFmtOperationCancelled           = "operation %s cancelled: %w"
	ErrFmtOperationFailedAfterRetries  = "operation %s failed after %d attempts: %w"
	ErrFmtOperationFailed              = "%s failed: %w"
	ErrFmtPanicCriticalOperationFailed = "critical operation %s failed: %v"
)
