package controlPlane

import (
	"blug/internal/pkg"
	"blug/internal/pkg/async"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"time"
)

var (
	flushCacheCh     = make(chan string, 20)
	getMachineInfoCh = make(chan string, 20)
	taskExecuted     = false
)

func InitDaemonSet() {
	go flushCacheProducer()
	go getMachineInfoProducer()
	if !taskExecuted {
		select {
		// 用10s 模拟系统低负载情况
		case <-time.After(time.Second * 10):
			flushCacheCh <- pkg.TASK_FLUSH_CACHE
			getMachineInfoCh <- pkg.TASK_GET_MACHINE_INFO
			taskExecuted = true
		}
	}
}

func flushCacheProducer() {
	for {
		select {
		case <-flushCacheCh:
			payload, err := json.Marshal(async.TaskPayload{
				TaskId:   uuid.New(),
				TaskName: pkg.TASK_FLUSH_CACHE,
				Status:   pkg.STATUS_INITIALIZING,
			})
			if err != nil {
				log.Error(err)
				continue
			}
			task := asynq.NewTask(pkg.TASK_FLUSH_CACHE, payload)
			info, err := async.GetAsynqClient().Enqueue(task, asynq.Queue(pkg.DEFAULT_QUEUE))
			if err != nil {
				log.Fatal(err)
			} else {
				log.Infof("Task enqueued successfully with ID: %s, Queue: %s, ETA: %s", info.ID, info.Queue, info.NextProcessAt)
			}
		}
	}
}

func getMachineInfoProducer() {
	for {
		select {
		case <-getMachineInfoCh:
			payload, err := json.Marshal(async.TaskPayload{
				TaskId:   uuid.New(),
				TaskName: pkg.TASK_GET_MACHINE_INFO,
				Status:   pkg.STATUS_INITIALIZING,
			})
			if err != nil {
				log.Error(err)
				continue
			}
			task := asynq.NewTask(pkg.TASK_GET_MACHINE_INFO, payload)
			info, err := async.GetAsynqClient().Enqueue(task)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Infof("Task enqueued successfully with ID: %s, Queue: %s, ETA: %s", info.ID, info.Queue, info.NextProcessAt)
			}
		}
	}
}

//func CancelJob()
