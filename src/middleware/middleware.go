/*
 * @Author: charley zhu
 * @Date: 2023-10-11 10:11:25
 * @LastEditTime: 2023-10-15 12:54:49
 * @LastEditors: charley zhu
 * @Description:
 */
package middleware

import (
	"global-backend/src/database"
	"global-backend/src/jwt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

/**
 * @description: middleware for jwt auth
 * @return {*}
 */
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if in debug mode, not check auth
		if os.Getenv("GIN_MODE") == "debug" {
			c.Next()
			return
		}

		// not check auth path
		if c.Request.URL.Path == "/user/authorize" {
			c.Next()
			return
		}

		// auth must in header
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "auth is empty",
			})
			c.Abort()
			return
		}

		// check token start with "Bearer"
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "auth format error",
			})
			c.Abort()
			return
		}

		// parse token
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("username", mc.UserName)
		c.Set("user_id", mc.UserId)

		c.Next()
	}
}

/**
 * @description: middleware for rate limit, every ip can only request 10 times per second
 * @return {*}
 */

var limiter = redis_rate.NewLimiter(database.GetRdb())
var limit = redis_rate.PerSecond(10)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := limiter.Allow(c.Request.Context(), c.ClientIP(), limit)
		if err != nil {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": 2006,
				"msg":  "rate limit",
			})
			c.Abort()
			return
		}

		if res.Allowed == 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": 2006,
				"msg":  "rate limit",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
