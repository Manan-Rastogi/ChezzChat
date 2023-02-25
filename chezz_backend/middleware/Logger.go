package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Logger(logFile *lumberjack.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the start time of the request
		start := time.Now()

		// Process the request
		ctx.Next()

		// Get the end time of the request
		end := time.Now()

		fmt.Fprintf(logFile, "[%s] \"%s %s %s\" %d %d \"%s\" %v\n",
			end.Format("2006-01-02 15:04:05.999999"),
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Request.Proto,
			ctx.Writer.Status(),
			ctx.Writer.Size(),
			ctx.Request.Referer(),
			end.Sub(start),
		)
	}
}
