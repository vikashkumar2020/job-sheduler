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

	for {
		jobs := *store.GetStoreInstance().GetStore()
		fmt.Println("job running ",len(jobs))
		if len(jobs) > 0 {

			foundPending := false
			for _, job := range jobs {
				if job.Status == "Pending" {
					store.GetStoreInstance().SaveJob(job.ID,"Running")
					foundPending = true
					fmt.Printf("Job %s started\n", job.Name)
					jobChannel <- "update"
					// Execute the job
					
					time.Sleep(job.Duration*time.Nanosecond)
					// Update job status to "Completed"
					job.Status = "Completed"
					job.UpdatedAt = time.Now()
					store.GetStoreInstance().SaveJob(job.ID,"Completed")
					fmt.Printf("Job %s completed\n", job.Name)

					// Send the updated job status through the channel
					jobChannel <- "update"

					// Exit the loop after handling one job
					fmt.Println("jobs ",jobs)
					time.Sleep(3*time.Second)
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
