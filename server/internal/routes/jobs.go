package routes

import (
	"job-sheduler/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterJobsRoutes(router *gin.RouterGroup) {
	router.GET("/jobs", controller.GetAllJobs)
	router.POST("/job", controller.CreateJob)
}