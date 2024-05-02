package store

import (
	model "job-sheduler/internal/model/entity"
	"time"

	"github.com/google/uuid"
)

// store to have all the jobs details
type Store struct {
	storeInstance *[]model.Job
}

// singelton store instance
var storeInstance *Store

// new store provider
func (store *Store) NewStore() {
	store.storeInstance = &[]model.Job{}
}

// function to follow singleton pattern to get store
func GetStoreInstance() *Store {
	if storeInstance == nil {
		storeInstance = &Store{}
	}
	return storeInstance
}

// getter function to get the store instance
func (store *Store) GetStore() *[]model.Job {
	return store.storeInstance
}

// function to create a new job
func (s *Store) CreateJob(job model.Job) {
	*s.storeInstance = append(*s.storeInstance, job)
}

// function to save the updated status of the jobs
func (s *Store) SaveJob(jobID uuid.UUID, status string) {
	jobs := *s.storeInstance
	for i := range jobs {
		if jobs[i].ID == jobID {
			jobs[i].Status = status
			jobs[i].UpdatedAt = time.Now()
			break
		}
	}
}
