/*
 * @Author: charley zhu
 * @Date: 2023-10-22 11:08:26
 * @LastEditTime: 2023-10-22 11:16:15
 * @LastEditors: charley zhu
 * @Description:
 */
package tool

import (
	"sync"
	"sync/atomic"
	"time"

	"global-backend/src/logger"
)

func WriteSomething(name string, c *sync.Cond, done *bool) {
	logger.Logger.Info("writeSomething start")
	time.Sleep(3 * time.Second)
	c.L.Lock()
	defer c.L.Unlock()
	*done = true
	c.Broadcast()
}

func ReadSomething(c *sync.Cond, done *bool) {
	logger.Logger.Info("readSomething start")
	c.L.Lock()
	defer c.L.Unlock()
	for !*done {
		c.Wait()
	}
	logger.Logger.Info("readSomething end")
}

func CreateBuffer(n *int32) interface{} {
	atomic.AddInt32(n, 1)
	buffer := make([]byte, 1024)
	return &buffer
}
