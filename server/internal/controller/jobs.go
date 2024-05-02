package controller

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"job-sheduler/internal/infra/store"
	socket "job-sheduler/internal/infra/websocket"
	model "job-sheduler/internal/model/entity"
	types "job-sheduler/internal/model/types"
	"job-sheduler/internal/service"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	jobMutex sync.Mutex // Mutex to synchronize access to jobs
)

func GetAllJobs(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	pool := socket.GetPoolInstance()

	defer conn.Close()
	// whenever recieve new update in jobstatus send all jobs to the clients
	client := &socket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client

	client.Read()

}

func CreateJob(c *gin.Context) {

	// get the request body
	var requestBody types.JobRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Invalid Request Body",
				"status":  "ERROR",
			},
		)
		return
	}

	// get the store instance
	storeData := store.GetStoreInstance()

	// create the result
	job := model.Job{
		ID:        uuid.New(),
		Name:      requestBody.Name,
		Duration:  time.Duration(requestBody.Time * uint64(time.Second)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "Pending",
	}

	// Acquire the lock
	jobMutex.Lock()
	defer jobMutex.Unlock()

	// save the job in the store
	storeData.CreateJob(job)
	service.SJFSchedule()
	pool := socket.GetPoolInstance()
	pool.Broadcast <- "update"
	c.JSON(
		200,
		gin.H{
			"message": "Job Added Successfully",
			"status":  "SUCCESS",
			"data":    *store.GetStoreInstance().GetStore(),
		},
	)
}
