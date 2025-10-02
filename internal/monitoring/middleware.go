package monitoring

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// MetricsMiddleware is a Gin middleware that records Prometheus metrics for each request.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Start a timer right when the request comes in.
		start := time.Now()

		// 2. Process the request.
		// c.Next() passes the request along to the actual handler (like CreateUdhaar).
		c.Next()

		// 3. After the handler has finished, the code below will run.
		// Calculate the request duration.
		duration := time.Since(start).Seconds()
		// Get the final HTTP status code (e.g., 200, 404, 500).
		statusCode := strconv.Itoa(c.Writer.Status())
		// Get the route path (e.g., "/api/v1/udhaars/:id").
		path := c.FullPath() 

		// 4. Update our metrics from the metrics.go file.
		// Increment the total request counter.
		HttpRequestsTotal.WithLabelValues(c.Request.Method, path, statusCode).Inc()
		// Record the duration in the histogram.
		HttpRequestDuration.WithLabelValues(c.Request.Method, path, statusCode).Observe(duration)
	}
}
