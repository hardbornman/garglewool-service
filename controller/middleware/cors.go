package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")//允许访问所有域
		//c.Writer.Header().Set("Access-Control-Allow-Headers","*")//header的类型
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*, accept, content-type, Authorization, authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH , HEAD, DELETE")
		if c.Request.Method == "OPTIONS" {
			//c.AbortWithStatus(204)   //放行预检
			/*请求方法不是GET/HEAD/POST
			POST请求的Content-Type并非application/x-www-form-urlencoded, multipart/form-data, 或text/plain
			请求设置了自定义的header字段*/
			return
		}
		c.Next()
	}
}
