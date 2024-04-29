/*
 * @Author: charley zhu
 * @Date: 2023-10-10 12:39:07
 * @LastEditTime: 2023-10-13 07:09:51
 * @LastEditors: charley zhu
 * @Description:
 */
package admin

import "github.com/gin-gonic/gin"

func AdminHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "admin hello",
	})
}
