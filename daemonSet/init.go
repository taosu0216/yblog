package daemonSet

import (
	"blug/daemonSet/controlPlane"
	"blug/daemonSet/workerNode"
)

func InitDaemonSet() {
	controlPlane.InitDaemonSet()
	workerNode.InitDaemonSet()
}
