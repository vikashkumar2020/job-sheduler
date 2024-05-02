package store

import (
	model "job-sheduler/internal/model/entity"
	"time"

	"github.com/google/uuid"
)
type Store struct {
	storeInstance *[]model.Job
}

var storeInstance *Store

func (store *Store) NewStore() {
	store.storeInstance = &[]model.Job{}
}

func GetStoreInstance() *Store {
	if storeInstance == nil {
		storeInstance = &Store{}
	}
	return storeInstance
}

func (store *Store) GetStore() *[]model.Job {
	return store.storeInstance
}

func (s *Store) CreateJob(job model.Job) {
	*s.storeInstance = append(*s.storeInstance, job)
}

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
