package controlPlane

import (
	"blug/internal/data/ent/task"
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
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
		obj.Infof("|| {controlPlane retry func}  task successfully: %s , [%s]", tp.TaskName, tp.TaskId)
	}
	return nil
}

func (obj controlPlaneObj) flushCacheRetryFunc(ctx context.Context, tp async.TaskPayload) error {
	obj.Infof("|| {controlPlane retry func} start task: %s , [%s]", tp.TaskName, tp.TaskId)
	i := 1
	for i = 1; i <= obj.maxRetries; i++ {
		obj.Infof("|| {controlPlane retry func} task: %s, [%s], retry: %d", tp.TaskName, tp.TaskId, i)
		err := obj.srv.RedisClient.Del(obj.ctx, pkg.ArticleListKey).Err()
		if err != nil {
			errStr := fmt.Sprintf("|| {controlPlane retry func} flush cache func %s [delete list key] err is: %v", tp.TaskId, err.Error())
			obj.Error(errStr)
			continue
		}
		obj.Infof("|| {controlPlane retry func} flush cache func [delete list key] [%s] succeed", tp.TaskId)

		err = obj.srv.RedisClient.Del(obj.ctx, pkg.ArticleMapKey).Err()
		if err != nil {
			errStr := fmt.Sprintf("|| {controlPlane retry func} flush cache func %s [delete map key] err is: %v", tp.TaskId, err.Error())
			obj.Error(errStr)
			continue
		}
		succStr := fmt.Sprintf("|| {controlPlane retry func} flush cache func [delete map key] [%s] succeed", tp.TaskId)
		obj.Info(succStr)

		tp.Status = pkg.STATUS_SUCCESS
		err = obj.srv.Pg.Task.Update().SetStatus(tp.Status).SetFinishTime(pkg.NowTimeStr()).Where(task.TaskID(tp.TaskId)).Exec(obj.ctx)
		if err != nil {
			errStr := fmt.Sprintf("|| {controlPlane retry func} flush cache func %s [update task status to db] err is: %v", tp.TaskId, err.Error())
			obj.Error(errStr)
			continue
		}
		succStr = fmt.Sprintf("|| {controlPlane retry func} flush cache func [%s] [update task status to db] succeed", tp.TaskId)
		obj.Info(succStr)
		break
	}
	if i == obj.maxRetries {
		errStr := fmt.Sprintf("|| {controlPlane retry func}  task failed: %s, [%s]", tp.TaskName, tp.TaskId)
		return errors.New(errStr)
	}
	return nil
}
