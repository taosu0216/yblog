package workerNode

import (
	"blug/internal/data/ent/task"
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"blug/internal/pkg/loggers"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
)

type workerNodeObj struct {
	srv       async.AsynqServer
	client    async.AsynqClient
	inspector async.AsynqInspector
	ctx       context.Context
	*loggers.ZapLogger
}

var workerNode workerNodeObj

func InitDaemonSet(ctx context.Context, skip int) {
	workerNode = workerNodeObj{}
	workerNode.srv = async.GetAsynqServer()
	workerNode.client = async.GetAsynqClient()
	workerNode.inspector = async.GetAsyncInspector()
	workerNode.ZapLogger = loggers.InitLog(pkg.GetRootLocation(), skip)
	workerNode.ctx = ctx
	workerNode.Info("{workerNode} Task log: start init server...")
	if err := workerNode.srv.AsynqDefaultServer.Run(asynq.HandlerFunc(workerNode.handler)); err != nil {
		workerNode.Error("{workerNode} Task log: run handler err is: " + err.Error())
	}
}

func (obj workerNodeObj) handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case pkg.TASK_FLUSH_CACHE:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			errStr := fmt.Sprintf("|| {workerNode} parse %s payload str to struct err is: %v", t.Type(), err)
			return errors.New(errStr)
		}
		err := obj.flushCacheFunc(obj.ctx, tp)
		if err != nil {
			return err
		}
	case pkg.TASK_GET_MACHINE_INFO:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			errStr := fmt.Sprintf("|| {workerNode} parse %s payload str to struct err is: %v", t.Type(), err)
			return errors.New(errStr)
		}
		err := obj.getMachineInfoFunc(obj.ctx, tp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj workerNodeObj) flushCacheFunc(ctx context.Context, tp async.TaskPayload) error {
	err := workerNode.srv.RedisClient.Del(obj.ctx, pkg.ArticleListKey).Err()
	if err != nil {
		errStr := fmt.Sprintf("|| {workerNode} flush cache func %s [delete list key] err is: %v", tp.TaskId, err.Error())
		obj.Error(errStr)

		tp.Status = pkg.STATUS_FAIL
		tp.Reason = err.Error()
		payload, err2 := json.Marshal(tp)
		if err2 != nil {
			errStr = fmt.Sprintf("|| {workerNode} flush cache func %s [delete list key] || obj to struct err is: %v", tp.TaskId, err.Error())
			return errors.New(errStr)
		}

		// todo: 死信队列接收失败任务及后续logic开发
		_, err3 := workerNode.client.AsynqClient.Enqueue(asynq.NewTask(tp.TaskType, payload), asynq.Queue(pkg.DEATH_QUEUE))
		if err3 != nil {
			errStr = fmt.Sprintf("|| {workerNode} flush cache func %s [delete list key] || failed job to death queue err is: %v", tp.TaskId, err3.Error())
			return errors.New(errStr)
		}

		return err
	}
	obj.Infof("|| {workerNode} flush cache func [delete list key] [%s] succeed", tp.TaskId)

	err = workerNode.srv.RedisClient.Del(obj.ctx, pkg.ArticleMapKey).Err()
	if err != nil {
		errStr := fmt.Sprintf("|| {workerNode} flush cache func %s [delete map key] err is: %v", tp.TaskId, err.Error())
		obj.Error(errStr)

		tp.Status = pkg.STATUS_FAIL
		tp.Reason = err.Error()
		payload, err2 := json.Marshal(tp)
		if err2 != nil {
			errStr = fmt.Sprintf("|| {workerNode} flush cache func %s [delete map key] || obj to struct err is: %v", tp.TaskId, err.Error())
			return errors.New(errStr)
		}

		// todo: 死信队列接收失败任务及后续logic开发
		_, err3 := workerNode.client.AsynqClient.Enqueue(asynq.NewTask(tp.TaskType, payload), asynq.Queue(pkg.DEATH_QUEUE))
		if err3 != nil {
			errStr = fmt.Sprintf("|| {workerNode} flush cache func %s [delete map key] || failed job to death queue err is: %v", tp.TaskId, err3.Error())
			return errors.New(errStr)
		}

		return err
	}
	succStr := fmt.Sprintf("|| {workerNode} flush cache func [delete map key] [%s] succeed", tp.TaskId)
	obj.Info(succStr)

	tp.Status = pkg.STATUS_SUCCESS
	err = workerNode.srv.Pg.Task.Update().SetStatus(tp.Status).SetFinishTime(pkg.NowTimeStr()).Where(task.TaskID(tp.TaskId)).Exec(obj.ctx)
	if err != nil {
		errStr := fmt.Sprintf("|| {workerNode} flush cache func %s [update task status to db] err is: %v", tp.TaskId, err.Error())
		return errors.New(errStr)
	}
	succStr = fmt.Sprintf("|| {workerNode} flush cache func [%s] [update task status to db] succeed", tp.TaskId)
	obj.Info(succStr)

	return nil
}

// todo: getMachineInfoFunc 任务logic
func (obj workerNodeObj) getMachineInfoFunc(ctx context.Context, tp async.TaskPayload) error {
	return nil
}
