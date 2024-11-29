package daemonSet

import (
	"blug/daemonSet/controlPlane"
	"blug/daemonSet/workerNode"
	"context"
)

func InitDaemonSet(ctx context.Context, isControlPlane bool, skip, maxRetries int) {
	if isControlPlane {
		controlPlane.InitDaemonSet(ctx, skip, maxRetries)
	} else {
		workerNode.InitDaemonSet(ctx, skip)
	}
}
