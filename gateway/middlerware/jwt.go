package middleware

import (
	"github.com/gin-gonic/gin"
	"grpc-todoList/pkg/res"
	"grpc-todoList/pkg/util"
	"time"
)

func JwT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(400, res.Response{
				Status: 400,
				Msg:    "token为空",
			})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(400, res.Response{
				Status: 400,
				Msg:    "token解析失败",
			})
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(400, res.Response{
				Status: 400,
				Msg:    "token过期",
			})
			c.Abort()
			return
		}
		//fmt.Println(time.Now().Unix())
		//fmt.Println(claims.ExpiresAt)
		c.Next()

	}
}
