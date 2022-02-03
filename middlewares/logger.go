package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %s %d \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Latency,
			params.Path,
			params.StatusCode)
	})
}

func SetUpLogOutPut() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

}
