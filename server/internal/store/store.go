package store

import (
	model "job-sheduler/internal/model/entity"
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

func (s *Store) SaveJob(job model.Job) {
	*s.storeInstance = append(*s.storeInstance, job)
}


