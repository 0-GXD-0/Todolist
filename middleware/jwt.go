package middleware

import (
	"demo1/pkg/utils"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		var msg string
		token := c.GetHeader("Authorization")
		// 去除 Bearer 前缀
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
		log.Println("Processed token:", token)
		if token == "" {
			code = 404
			msg = "Token 缺失"
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				log.Printf("错误： %v, token: %s", err, token)
				code = 403
				msg = "无效的 Token"
			} else if claim == nil {
				log.Printf("Token解析失败，claim 为 nil")
				code = 403 // 无权限
				msg = "Token 解析失败"
			} else if time.Now().Unix() > claim.ExpiresAt {
				log.Printf("Token解析成功，claim: %+v", claim)
				code = 401 //token无效
				msg = "Token 已过期"
			}
		}
		if code != 200 {
			log.Printf("响应状态码: %d, 消息: %s", code, msg)
			c.JSON(200, gin.H{
				"status": code,
				"msg":    msg,
			})
			c.Abort()
			return
		}
		log.Println("Before c.Next()")
		c.Next()
		log.Println("After c.Next()")
		log.Println("JWT middleware completed successfully")
	}
}
