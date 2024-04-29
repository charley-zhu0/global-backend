/*
 * @Author: charley zhu
 * @Date: 2023-10-10 12:39:07
 * @LastEditTime: 2023-10-23 12:51:06
 * @LastEditors: charley zhu
 * @Description:
 */
package user

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"global-backend/src/jwt"
	"global-backend/src/logger"
	"global-backend/src/tool"
)

func UserHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user hello",
	})
}

func UserAuthorize(c *gin.Context) {
	parm := &AuthSchemaValidator{}
	if err := c.ShouldBindJSON(parm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 4000,
			"msg":  "invalid param",
		})
		return
	}
	username := parm.UserName
	password := parm.Password

	if username != "root" || password != "hello" {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 4030,
			"msg":  "invalid username or password",
		})
		return
	}

	// gen token
	token, err := jwt.GenToken(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 5000,
			"msg":  "gen token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": token,
	})
}

func UserSyncCond(c *gin.Context) {
	// test usage of sync.Cond
	// one wait for multiple use waitgroup
	// multiple wait for one use cond
	cond := &sync.Cond{L: &sync.Mutex{}}
	done := false
	go tool.ReadSomething(cond, &done)
	go tool.ReadSomething(cond, &done)
	go tool.ReadSomething(cond, &done)
	go tool.ReadSomething(cond, &done)
	tool.WriteSomething("hello", cond, &done)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
	})
}

func UserSyncPool(c *gin.Context) {
	// test usage of sync.Pool
	var count int32

	var myfunc = func() interface{} {
		return tool.CreateBuffer(&count)
	}

	bufferPool := &sync.Pool{
		New: myfunc,
	}

	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	logger.Logger.Info(string(count))
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
	})
}
