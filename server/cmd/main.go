package main

import (
	register "job-sheduler/internal/common/register"
	config "job-sheduler/internal/config"
	"job-sheduler/internal/service"
	store "job-sheduler/internal/store"
	utils "job-sheduler/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the config
	config.LoadEnv()
	utils.LogInfo("env loaded ")

	// Initialize the server
	serverConfig := config.NewServerConfig()
	utils.LogInfo("server config loaded")

	// initialize store
	store := store.GetStoreInstance()
	store.NewStore()
	utils.LogInfo("store initialized")
	router := gin.Default()
	register.Routes(router, serverConfig)
	
	go service.UpdateJobStatus(store.GetQueue())

	if err := router.Run(":" + serverConfig.Port); err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("server started")
	
}
