package daemonSet

import (
	"blug/daemonSet/controlPlane"
	"blug/daemonSet/workerNode"
	"context"
)

func InitDaemonSet(ctx context.Context, isControlPlane bool, skip int) {
	if isControlPlane {
		controlPlane.InitDaemonSet(ctx, skip)
	} else {
		workerNode.InitDaemonSet(ctx, skip)
	}
}
