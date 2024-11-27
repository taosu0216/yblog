package workerNode

import (
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

func InitDaemonSet() {
	if err := async.GetAsynqServer().Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case pkg.TASK_FLUSH_CACHE:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			return err
		}
		err := flushCacheFunc(ctx, tp)
		if err != nil {
			return err
		}
	case pkg.TASK_GET_MACHINE_INFO:
		var tp async.TaskPayload
		if err := json.Unmarshal(t.Payload(), &tp); err != nil {
			return err
		}
		err := getMachineInfoFunc(ctx, tp)
		if err != nil {
			return err
		}
	}
	return nil
}

func flushCacheFunc(ctx context.Context, tp async.TaskPayload) error {
	//todo: update the status to STATUS_RUNNING in db   	tp.Status = pkg.STATUS_RUNNING
	err := async.GetAsynqServer().RedisClient.Del(ctx, pkg.ArticleListKey).Err()
	if err != nil {
		log.Println(err)
		tp.Status = pkg.STATUS_FAIL
		tp.Reason = err.Error()
		payload, err2 := json.Marshal(tp)
		if err2 != nil {
			log.Println(err2)
			return err
		}
		// todo: 死信队列处理中入库
		_, err3 := async.GetAsynqClient().Enqueue(asynq.NewTask(pkg.TASK_FLUSH_CACHE, payload), asynq.Queue(pkg.DEFAULT_QUEUE))
		if err3 != nil {
			log.Println(err3)
			return err
		}
		return err
	} else {
		log.Println("delete list key succeed")
	}
	err = async.GetAsynqServer().RedisClient.Del(ctx, pkg.ArticleMapKey).Err()
	if err != nil {
		log.Println(err)
		tp.Status = pkg.STATUS_FAIL
		tp.Reason = err.Error()
		return err
	} else {
		log.Println("delete map key succeed")
	}
	tp.Status = pkg.STATUS_SUCCESS
	// todo: 入库
	return nil
}

// todo
func getMachineInfoFunc(ctx context.Context, tp async.TaskPayload) error {
	return nil
}
