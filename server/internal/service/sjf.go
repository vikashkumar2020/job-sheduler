package service

import (
	"fmt"
	"job-sheduler/internal/infra/store"
	"job-sheduler/internal/infra/websocket"
	"sort"
	"time"
)

func SJFSchedule() {

	jobs := *store.GetStoreInstance().GetStore()

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Duration < jobs[j].Duration
	})

}

func UpdateJobStatus(jobChannel chan string) {

	for {
		jobs := *store.GetStoreInstance().GetStore()
		fmt.Println("Total Jobs ", len(jobs))
		if len(jobs) > 0 {

			foundPending := false
			for _, job := range jobs {
				if job.Status == "Pending" {
					store.GetStoreInstance().SaveJob(job.ID, "Running")
					foundPending = true
					fmt.Printf("Job %s started\n", job.Name)
					if len(websocket.GetPoolInstance().Clients) > 0 {
						jobChannel <- fmt.Sprintf("%s started having status Running", job.Name)
					}

					time.Sleep(job.Duration)
					// Update job status to "Completed"
					job.Status = "Completed"
					job.UpdatedAt = time.Now()
					store.GetStoreInstance().SaveJob(job.ID, "Completed")
					fmt.Printf("Job %s completed\n", job.Name)

					// Send the updated job status through the channel
					if len(websocket.GetPoolInstance().Clients) > 0 {
						jobChannel <- fmt.Sprintf("%s started having status Completed", job.Name)
					}
					break
				}
			}

			if !foundPending {
				time.Sleep(1 * time.Second)
			}
		} else {
			time.Sleep(1 * time.Second) // Sleep if there are no jobs
		}
	}
}
