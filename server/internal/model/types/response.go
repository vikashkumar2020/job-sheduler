package types

import model "job-sheduler/internal/model/entity"

type JobResponse struct {
	Jobs []model.Job `json:"jobs"`
	Length int `length:"jobs"`
}