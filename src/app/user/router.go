/*
 * @Author: charley zhu
 * @Date: 2023-10-10 12:38:37
 * @LastEditTime: 2023-10-23 13:08:18
 * @LastEditors: charley zhu
 * @Description:
 */
package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/hello", UserHello)
		userGroup.POST("/authorize", UserAuthorize)
		userGroup.GET("/test/syncCond", UserSyncCond)
		userGroup.GET("/test/syncPool", UserSyncPool)
	}
}
