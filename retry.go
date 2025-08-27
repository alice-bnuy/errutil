package errutil

import (
	"alicemains/internal/shared/logger"
	"fmt"
	"time"
)

func (rm *RetryManager) ExecuteWithRetry(operation func() error, operationName string) error {
	var lastErr error

	for attempt := 0; attempt <= rm.maxRetries; attempt++ {
		if attempt > 0 {
			delay := rm.calculateDelay(attempt)
			logger.WithField("attempt", attempt+1).
				WithField("max_attempts", rm.maxRetries+1).
				WithField("operation", operationName).
				WithField("delay", delay).
				Infof("Attempt %d/%d for %s in %v", attempt+1, rm.maxRetries+1, operationName, delay)
			time.Sleep(delay)
		}

		err := operation()
		if err == nil {
			if attempt > 0 {
				logger.WithField("operation", operationName).
					WithField("attempts_used", attempt+1).
					Infof("Operation %s succeeded after %d attempts", operationName, attempt+1)
			}
			return nil
		}

		lastErr = err
		logger.WithField("operation", operationName).
			WithField("attempt", attempt+1).
			ErrorWithErr(fmt.Sprintf(ErrOnAttempt, attempt+1, operationName), err)
	}

	return fmt.Errorf(ErrOperationAttemptsFailed, operationName, rm.maxRetries+1, lastErr)
}

func (rm *RetryManager) calculateDelay(attempt int) time.Duration {
	delay := time.Duration(float64(rm.baseDelay) * float64(attempt) * rm.backoffFactor)
	if delay > rm.maxDelay {
		delay = rm.maxDelay
	}
	return delay
}

func (rm *RetryManager) SetMaxRetries(retries int) {
	rm.maxRetries = retries
}

func (rm *RetryManager) SetBaseDelay(delay time.Duration) {
	rm.baseDelay = delay
}

func (rm *RetryManager) SetMaxDelay(delay time.Duration) {
	rm.maxDelay = delay
}
