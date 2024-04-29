/*
 * @Author: charley zhu
 * @Date: 2023-10-10 12:30:33
 * @LastEditTime: 2023-10-15 07:55:48
 * @LastEditors: charley zhu
 * @Description:
 */
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"global-backend/src/lib"
	"global-backend/src/logger"
)

func main() {
	// get gin engine
	e, err := lib.SetupRouter()
	if err != nil {
		logger.Logger.Fatal("SetupRouter failed, err:", zap.Error(err))
		return
	}

	// graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// channel for getting signal
	quit := make(chan os.Signal, 1)
	// kill -> syscall.SIGTERM
	// kill -2  -> syscall.SIGINT (like Ctrl+C)
	// kill -9 -> syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit // wait for signal
	// create a context for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server Shutdown:", zap.Error(err))
	}
	logger.Logger.Info("Server exiting")
}
