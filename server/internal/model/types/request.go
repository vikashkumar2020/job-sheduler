package types

// Request format for submitting a new job
type JobRequest struct {
	Name string    `json:"name"`
	Time uint64    `json:"duration"`
}
