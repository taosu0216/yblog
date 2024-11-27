package async

import "github.com/google/uuid"

type TaskPayload struct {
	TaskId   uuid.UUID
	TaskName string
	Status   string
	Reason   string
}
