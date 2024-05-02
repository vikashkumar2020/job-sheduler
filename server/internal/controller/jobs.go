package controller

import (
	// "fmt"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	model "job-sheduler/internal/model/entity"
	types "job-sheduler/internal/model/types"
	"job-sheduler/internal/service"
	"job-sheduler/internal/store"
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

	jobStatus := store.GetStoreInstance().GetQueue()
	defer conn.Close();
	// whenever recieve new update in jobstatus send all jobs to the clients
	for {
		select {
		case jobs := <-jobStatus:
			// Send job updates to the client
			list := *store.GetStoreInstance().GetStore()
			fmt.Println("update",jobs)
			response := types.JobResponse{
				Jobs: list,
				Length: len(list),
			}
			if err := conn.WriteJSON(response); err != nil {
				fmt.Println("Error writing to WebSocket:", err)
				return
			}
		}
	}
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
	store := store.GetStoreInstance()

	// create the result
	job := model.Job{
		ID:        uuid.New(),
		Name:      requestBody.Name,
		Duration:  time.Duration(requestBody.Time*uint64(time.Millisecond)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "Pending",
	}

	// Acquire the lock
	jobMutex.Lock()
	defer jobMutex.Unlock()

	// save the job in the store
	store.CreateJob(job)
	service.SJFSchedule()

	c.JSON(
		200,
		gin.H{
			"message": "Job Added Successfully",
			"status":  "SUCCESS",
		},
	)
}
