package middlewares

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	logFilePath := "./"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println("Failed to create log file:", err)
	}
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		logMessage := fmt.Sprintf("%s - [%s] %s %s %d \n",
			params.ClientIP,
			params.TimeStamp,
			params.Method,
			params.Path,
			params.StatusCode,
		)

		// 將日誌訊息寫入到指定路徑的日誌檔案中
		if _, err := logFile.WriteString(logMessage); err != nil {
			fmt.Println("Failed to write log:", err)
		}

		return logMessage
	})
}
