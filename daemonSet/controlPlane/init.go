package controlPlane

import (
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"blug/internal/pkg/loggers"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

var (
	flushCacheCh     = make(chan async.TaskPayload, 20)
	getMachineInfoCh = make(chan async.TaskPayload, 20)
)

type controlPlaneObj struct {
	srv       async.AsynqServer
	client    async.AsynqClient
	inspector async.AsynqInspector
	*loggers.ZapLogger
}

var controlPlane controlPlaneObj

func InitDaemonSet(ctx context.Context) {
	controlPlane = controlPlaneObj{
		srv:       async.GetAsynqServer(),
		client:    async.GetAsynqClient(),
		inspector: async.GetAsyncInspector(),
		ZapLogger: loggers.InitLog(pkg.GetRootLocation()),
	}

	go controlPlane.flushCacheProducer(ctx)
	go controlPlane.getMachineInfoProducer(ctx)
	go initDeathQueue()

	select {
	// 用1h 模拟系统低负载情况
	case <-time.After(time.Hour * 1):
		flushCacheCh <- async.TaskPayload{
			TaskType: pkg.TASK_FLUSH_CACHE,
			TaskName: pkg.TASK_FLUSH_CACHE + "-1-" + pkg.NowTimeStr(),
			Status:   pkg.STATUS_INITIALIZING,
			Queue:    pkg.DEFAULT_QUEUE,
		}
		getMachineInfoCh <- async.TaskPayload{
			TaskType: pkg.TASK_GET_MACHINE_INFO,
			TaskName: pkg.TASK_GET_MACHINE_INFO + "-1-" + pkg.NowTimeStr(),
			Status:   pkg.STATUS_INITIALIZING,
			Queue:    pkg.DEFAULT_QUEUE,
		}
	}
}

func (obj controlPlaneObj) flushCacheProducer(ctx context.Context) {
	for {
		select {
		case job := <-flushCacheCh:
			payload, err := json.Marshal(job)
			if err != nil {
				obj.Error("Task log: obj to str err is:", err)
				continue
			}

			task := asynq.NewTask(job.TaskType, payload)

			infos, err := obj.client.AsynqClient.Enqueue(task, asynq.Queue(job.Queue))
			if err != nil {
				obj.Errorf("Task log: %v to queue failed,task type is: %s, err is: %s", infos.ID, infos.Type, err)
				continue
			} else {
				obj.Infof("Task log: %v to queue:%s successfully,task type is: %s, ETA is: %s", infos.ID, infos.Queue, infos.Type, infos.NextProcessAt)
			}

			err = obj.client.Pg.Task.Create().SetTaskID(infos.ID).SetTaskName(job.TaskName).SetStatus(job.Status).Exec(ctx)
			if err != nil {
				obj.Errorf("Task log: %v to DB failed,err is: %s", infos.ID, err)
			}
		}
	}
}

func (obj controlPlaneObj) getMachineInfoProducer(ctx context.Context) {
	for {
		select {
		case job := <-getMachineInfoCh:
			payload, err := json.Marshal(job)
			if err != nil {
				obj.Error("Task log: obj to str err is:", err)
				continue
			}

			task := asynq.NewTask(job.TaskType, payload)

			infos, err := obj.client.AsynqClient.Enqueue(task, asynq.Queue(job.Queue))
			if err != nil {
				obj.Errorf("Task log: %v to queue failed,task type is: %s, err is: %s", infos.ID, infos.Type, err)
				continue
			} else {
				obj.Infof("Task log: %v to queue:%s successfully,task type is: %s, ETA is: %s", infos.ID, infos.Queue, infos.Type, infos.NextProcessAt)
			}

			err = obj.client.Pg.Task.Create().SetTaskID(infos.ID).SetTaskName(job.TaskName).SetStatus(job.Status).Exec(ctx)
			if err != nil {
				obj.Errorf("Task log: %v to DB failed,err is: %s", infos.ID, err)
			}
		}
	}
}

//func CancelJob()
