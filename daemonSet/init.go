package daemonSet

import (
	"blug/daemonSet/controlPlane"
	"blug/daemonSet/workerNode"
	"context"
)

func InitDaemonSet(ctx context.Context) {
	controlPlane.InitDaemonSet(ctx)
	workerNode.InitDaemonSet(ctx)
}
