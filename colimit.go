package colimit

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

// New create a limiter middleware
func New(limit int64) echo.MiddlewareFunc {
	// if limit <= 0, returns a empty middleware
	if limit <= 0 {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return next
		}
	}
	// the counter
	var count int64
	// the middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// check limit
			defer atomic.AddInt64(&count, -1)
			if atomic.AddInt64(&count, 1) > limit {
				return c.String(http.StatusInsufficientStorage, "concurrency limit exceeded")
			}
			// invoke next handler
			return next(c)
		}
	}
}
