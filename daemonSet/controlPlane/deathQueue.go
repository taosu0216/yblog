package controlPlane

import (
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

func (obj controlPlaneObj) initDeathQueue() {
	obj.Info("{controlPlane} Task log: start init Death Queue...")
	if err := obj.srv.AsynqDeathServer.Run(asynq.HandlerFunc(obj.handler)); err != nil {
		obj.Error("{controlPlane} Task log: run Death Queue handler err is: " + err.Error())
	}
}

func (obj controlPlaneObj) handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case pkg.TASK_FLUSH_CACHE:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			errStr := fmt.Sprintf("|| {controlPlane retry func} parse %s payload str to struct err is: %v", t.Type(), err)
			return errors.New(errStr)
		}
		err := obj.flushCacheRetryFunc(ctx, tp)
		if err != nil {
			errStr := fmt.Sprintf("|| {controlPlane retry func} flush cache err is: %v", err.Error())
			return errors.New(errStr)
		}
	}
	return nil
}

func (obj controlPlaneObj) flushCacheRetryFunc(ctx context.Context, tp async.TaskPayload) error {
	log.Println("Task log: flush cache retry func start")
	return nil
}
