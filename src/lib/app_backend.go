/*
 * @Author: charley zhu
 * @Date: 2023-10-13 06:43:06
 * @LastEditTime: 2023-10-15 09:34:45
 * @LastEditors: charley zhu
 * @Description:
 */
package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"global-backend/src/app/admin"
	"global-backend/src/app/user"
	"global-backend/src/config"
	"global-backend/src/logger"
	"global-backend/src/middleware"
)

/**
 * @description: register middleware and routers
 * @return {*}
 */
func SetupRouter() (*gin.Engine, error) {
	// load env from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("use default env")
	}

	// load config from json file
	// var env_path = os.Getenv("ENV_PATH")
	var env_path = "config.json"

	// init config to GlobalConfig
	if err := config.Init(env_path); err != nil {
		fmt.Println("load config failed, err:", err)
		return nil, err
	}

	// init logger
	if err := logger.InitLogger(config.GlobalConfig.LogConfig); err != nil {
		return nil, err
	}

	// set gin mode
	gin.SetMode(config.GlobalConfig.Mode)

	// create gin engine, if use gin.Default(), will use logger in gin
	e := gin.New()

	// use logger middleware
	e.Use(logger.GinLogger(), logger.GinRecovery(true))

	// register others middleware ....
	e.Use(middleware.JWTAuthMiddleware())
	e.Use(middleware.RateLimitMiddleware())

	// load routers
	routerList := []func(*gin.Engine){
		user.Routers,
		admin.Routers,
	}

	for _, router := range routerList {
		router(e)
	}

	return e, nil
}
