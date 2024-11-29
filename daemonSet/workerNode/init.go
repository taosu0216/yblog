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
	"log"
)

type workerNodeObj struct {
	srv       async.AsynqServer
	client    async.AsynqClient
	inspector async.AsynqInspector
	*loggers.ZapLogger
}

var workerNode workerNodeObj

func InitDaemonSet() {
	workerNode = workerNodeObj{}
	workerNode.srv = async.GetAsynqServer()
	workerNode.client = async.GetAsynqClient()
	workerNode.inspector = async.GetAsyncInspector()
	workerNode.ZapLogger = loggers.InitLog(pkg.GetRootLocation())
	if err := workerNode.srv.AsynqDefaultServer.Run(asynq.HandlerFunc(workerNode.handler)); err != nil {
		workerNode.Error("Task log: workerNode run handler err is: " + err.Error())
	} else {
		workerNode.Info("Task log: workerNode run handler succeed")
	}
}

func (obj workerNodeObj) handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case pkg.TASK_FLUSH_CACHE:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			errStr := fmt.Sprintf("|| parse %s payload str to struct err is: %v", t.Type(), err)
			return errors.New(errStr)
		}
		err := obj.flushCacheFunc(ctx, tp)
		if err != nil {
			return err
		}
	case pkg.TASK_GET_MACHINE_INFO:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			errStr := fmt.Sprintf("|| parse %s payload str to struct err is: %v", t.Type(), err)
			return errors.New(errStr)
		}
		err := obj.getMachineInfoFunc(ctx, tp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj workerNodeObj) flushCacheFunc(ctx context.Context, tp async.TaskPayload) error {

	err := workerNode.srv.RedisClient.Del(ctx, pkg.ArticleListKey).Err()
	if err != nil {
		obj.Error("|| flush cache func delete list key err is: " + err.Error())

		tp.Status = pkg.STATUS_FAIL
		tp.Reason = err.Error()
		payload, err2 := json.Marshal(tp)
		if err2 != nil {
			return errors.New("|| flush cache func delete list key || obj to struct err is: " + err2.Error())
		}

		// todo: 死信队列接收失败任务及后续logic开发
		_, err3 := workerNode.client.AsynqClient.Enqueue(asynq.NewTask(tp.TaskType, payload), asynq.Queue(pkg.DEATH_QUEUE))
		if err3 != nil {
			return errors.New("|| flush cache failed job to death queue err is: " + err3.Error())
		}

		return err
	} else {
		obj.Info("delete list key succeed")
	}

	err = workerNode.srv.RedisClient.Del(ctx, pkg.ArticleMapKey).Err()
	if err != nil {
		obj.Error("|| flush cache func delete map key err is: " + err.Error())

		tp.Status = pkg.STATUS_FAIL
		tp.Reason = err.Error()
		payload, err2 := json.Marshal(tp)
		if err2 != nil {
			return errors.New("|| flush cache func to struct err is: " + err2.Error())
		}

		// todo: 死信队列接收失败任务及后续logic开发
		_, err3 := workerNode.client.AsynqClient.Enqueue(asynq.NewTask(tp.TaskType, payload), asynq.Queue(pkg.DEATH_QUEUE))
		if err3 != nil {
			return errors.New("|| flush cache failed job to death queue err is: " + err3.Error())
		}

		return err
	} else {
		log.Println("delete map key succeed")
	}
	tp.Status = pkg.STATUS_SUCCESS
	err = workerNode.srv.Pg.Task.Update().SetStatus(tp.Status).SetFinishTime(pkg.NowTimeStr()).Where(task.TaskID(tp.TaskId)).Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// todo: getMachineInfoFunc 任务logic
func (obj workerNodeObj) getMachineInfoFunc(ctx context.Context, tp async.TaskPayload) error {
	return nil
}
