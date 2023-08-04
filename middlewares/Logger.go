package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string { //將func帶入LoggerFormatterParams 並回傳字串
		return fmt.Sprintf("%s - [%s] %s %s %d \n",
			params.ClientIP,   //IP
			params.TimeStamp,  //時間
			params.Method,     //方法
			params.Path,       //路徑
			params.StatusCode, //http StatusCode
		)
	})
}
