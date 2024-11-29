package async

type TaskPayload struct {
	TaskId   string `json:"task_id"`
	TaskType string `json:"task_type"`
	TaskName string `json:"task_name"`
	Queue    string `json:"queue"`
	Retry    int    `json:"retry"`
	Status   string
	Reason   string
}
