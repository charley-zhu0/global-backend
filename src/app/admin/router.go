/*
 * @Author: charley zhu
 * @Date: 2023-10-10 12:38:37
 * @LastEditTime: 2023-10-11 07:50:31
 * @LastEditors: charley_zhu@trendmicro.com
 * @Description:
 */
package admin

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	adminGroup := e.Group("/admin")
	{
		adminGroup.GET("/hello", AdminHello)
	}
}
