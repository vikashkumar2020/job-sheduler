package main

import (
	register "job-sheduler/internal/common/register"
	config "job-sheduler/internal/config"
	store "job-sheduler/internal/infra/store"
	"job-sheduler/internal/infra/websocket"
	"job-sheduler/internal/service"
	utils "job-sheduler/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the config
	config.LoadEnv()
	utils.LogInfo("env loaded for configuration")

	// Initialize the server
	serverConfig := config.NewServerConfig()
	utils.LogInfo("server config loaded")

	// initialize store
	store := store.GetStoreInstance()
	store.NewStore()
	utils.LogInfo("store initialized")

	pool := websocket.GetPoolInstance()
	pool.NewPool()
	utils.LogInfo("websockets pool initialized")

	router := gin.Default()
	register.Routes(router, serverConfig)
	utils.LogInfo("rigestered routes")

	go service.UpdateJobStatus(pool.Broadcast)
	utils.LogInfo("job routine is started")

	go pool.Start()
	utils.LogInfo("websocket listner is started")

	if err := router.Run(":" + serverConfig.Port); err != nil {
		utils.LogFatal(err)
	}

	utils.LogInfo("server started")

}
