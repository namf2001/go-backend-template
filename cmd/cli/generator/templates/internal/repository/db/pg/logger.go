package pg

import (
	"context"
	"time"

	"github.com/namf2001/go-backend-template/internal/pkg/logger"
)

// LogQuery logs the start and end of a SQL query execution.
// Use this in repository methods for debugging purposes.
func LogQuery(ctx context.Context, operation string, query string) func() {
	start := time.Now()
	logger.DEBUG.Printf("DB %s START: %s", operation, query)
	return func() {
		logger.DEBUG.Printf("DB %s END: %s (took %dms)", operation, query, time.Since(start).Milliseconds())
	}
}
