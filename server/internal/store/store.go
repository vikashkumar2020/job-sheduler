package store

import (
	model "job-sheduler/internal/model/entity"
	"time"

	"github.com/google/uuid"
)


type Store struct {
	storeInstance *[]model.Job
	jobUpdates chan string
}

var databaseInstance *Store


func (store *Store) NewStore() {
	store.storeInstance = &[]model.Job{}
	store.jobUpdates = make(chan string)
}

func GetStoreInstance() *Store {
	if databaseInstance == nil {
		databaseInstance = &Store{}
	}
	return databaseInstance
}

func (store *Store) GetStore() *[]model.Job {
	return store.storeInstance
}

func (store *Store) GetQueue() chan string{
	return store.jobUpdates
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