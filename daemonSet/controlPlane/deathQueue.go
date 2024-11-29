package controlPlane

import (
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

func initDeathQueue() {
	log.Println(controlPlane.srv.AsynqDeathServer.Run(asynq.HandlerFunc(controlPlane.handler)))
}

func (obj controlPlaneObj) handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case pkg.TASK_FLUSH_CACHE:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			return err
		}
		//err := flushCacheFunc(ctx, tp)
		//if err != nil {
		//	return err
		//}
	}
	return nil
}
