package controlPlane

import (
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"blug/internal/pkg/loggers"
	"context"
	"encoding/json"
	"github.com/google/uuid"
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
	ctx       context.Context
	*loggers.ZapLogger
}

var controlPlane controlPlaneObj

func InitDaemonSet(ctx context.Context, skip int) {
	controlPlane = controlPlaneObj{
		srv:       async.GetAsynqServer(),
		client:    async.GetAsynqClient(),
		inspector: async.GetAsyncInspector(),
		ZapLogger: loggers.InitLog(pkg.GetRootLocation(), skip),
		ctx:       ctx,
	}

	go controlPlane.flushCacheProducer()
	//go controlPlane.getMachineInfoProducer()
	go controlPlane.initDeathQueue()

	controlPlane.Info("{controlPlane} Task log: start init server...")
	for {
		select {
		// 用1h 模拟系统低负载情况
		case <-time.After(time.Second * 10):
			controlPlane.Info("{controlPlane} Task log: system low load, ready to produce job...")
			flushCacheCh <- async.TaskPayload{
				TaskId:   uuid.NewString(),
				TaskType: pkg.TASK_FLUSH_CACHE,
				TaskName: pkg.TASK_FLUSH_CACHE + "-" + pkg.NowTimeStr() + "--{1}",
				Status:   pkg.STATUS_INITIALIZING,
				Queue:    pkg.DEFAULT_QUEUE,
			}
			//getMachineInfoCh <- async.TaskPayload{
			//	TaskType: pkg.TASK_GET_MACHINE_INFO,
			//	TaskName: pkg.TASK_GET_MACHINE_INFO + "-1-" + pkg.NowTimeStr(),
			//	Status:   pkg.STATUS_INITIALIZING,
			//	Queue:    pkg.DEFAULT_QUEUE,
			//}
		}
	}
}

func (obj controlPlaneObj) flushCacheProducer() {
	for {
		select {
		case job := <-flushCacheCh:
			obj.Info("{controlPlane} Task log: start produce job...")
			payload, err := json.Marshal(job)
			if err != nil {
				obj.Error("{controlPlane} Task log: obj to str err is:", err)
				continue
			}

			task := asynq.NewTask(job.TaskType, payload)

			infos, err := obj.client.AsynqClient.Enqueue(task, asynq.Queue(job.Queue))
			if err != nil {
				obj.Errorf("{controlPlane} Task log: %v to queue failed,task type is: %s, err is: %s", job.TaskId, infos.Type, err)
				continue
			} else {
				obj.Infof("{controlPlane} Task log: %v to queue:%s successfully,task type is: %s, ETA is: %s", job.TaskId, infos.Queue, infos.Type, infos.NextProcessAt)
			}

			err = obj.client.Pg.Task.Create().SetTaskID(job.TaskId).SetTaskName(job.TaskName).SetStatus(job.Status).Exec(obj.ctx)
			if err != nil {
				obj.Errorf("{controlPlane} Task log: %v to db failed,err is: %s", job.TaskId, err)
			}
		}
	}
}

func (obj controlPlaneObj) getMachineInfoProducer() {
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
				obj.Errorf("Task log: %v to queue failed,task type is: %s, err is: %s", job.TaskId, infos.Type, err)
				continue
			} else {
				obj.Infof("Task log: %v to queue:%s successfully,task type is: %s, ETA is: %s", job.TaskId, infos.Queue, infos.Type, infos.NextProcessAt)
			}

			err = obj.client.Pg.Task.Create().SetTaskID(job.TaskId).SetTaskName(job.TaskName).SetStatus(job.Status).Exec(obj.ctx)
			if err != nil {
				obj.Errorf("Task log: %v to DB failed,err is: %s", job.TaskId, err)
			}
		}
	}
}

//func CancelJob()
