package even

import (
	"context"
	"log/slog"
)

func IsEven(ctx context.Context, number int) (bool, error) {
	// Extract request ID from context for logging
	var requestID string
	if id := ctx.Value("requestID"); id != nil {
		if idStr, ok := id.(string); ok {
			requestID = idStr
		}
	}

	slog.Debug("IsEven called", "request_id", requestID, "number", number)

	// TODO: Implement error handling for negative numbers
	if number%2 == 0 {
		slog.Debug("IsEven result", "request_id", requestID, "number", number, "is_even", true)
		return true, nil
	}
	slog.Debug("IsEven result", "request_id", requestID, "number", number, "is_even", false)
	return false, nil
}
