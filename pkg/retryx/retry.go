package retryx

import (
	"fmt"
	"time"
)

func Do(attempts int, baseDelay time.Duration, fn func(attempt int) error) error {
	var lastErr error

	for attempt := 1; attempt <= attempts; attempt++ {
		if err := fn(attempt); err != nil {
			lastErr = err
			if attempt == attempts {
				break
			}
			time.Sleep(time.Duration(attempt) * baseDelay)
			continue
		}

		return nil
	}

	return fmt.Errorf("all retries failed: %w", lastErr)
}
