package service

import (
	"fmt"
	"job-sheduler/internal/store"
	"sort"
	"time"
)

func SJFSchedule() {

	jobs := *store.GetStoreInstance().GetStore()
	
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Duration < jobs[j].Duration
	})

}

func UpdateJobStatus(jobChannel chan string)  {

	jobs := *store.GetStoreInstance().GetStore()
	
	for {
		if len(jobs) > 0 {
			for _, job := range jobs {
				if job.Status == "Pending" {
					// Update job status to "Running"
					job.UpdatedAt = time.Now()
					job.Status = "Running"
					fmt.Printf("Job %d started\n", job.ID)
					jobChannel <- "update"
					// Execute the job
					
					time.Sleep(job.Duration)
					// Update job status to "Completed"
					job.Status = "Completed"
					job.UpdatedAt = time.Now()
					fmt.Printf("Job %d completed\n", job.ID)

					// Send the updated job status through the channel
					jobChannel <- "update"

					// Exit the loop after handling one job
					break
				}
			}
		} else {
			time.Sleep(1 * time.Second) // Sleep if there are no jobs
		}
	}
}
