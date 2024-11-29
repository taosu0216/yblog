package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

func main() {
	// 获取初始的网络 I/O 统计
	initialIOCounters, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error getting initial network I/O counters:", err)
		return
	}

	// 等待一段时间
	time.Sleep(30 * time.Second)

	// 获取当前的网络 I/O 统计
	currentIOCounters, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("Error getting current network I/O counters:", err)
		return
	}

	// 计算网络 I/O 占用情况
	for i, current := range currentIOCounters {
		initial := initialIOCounters[i]
		fmt.Printf("Interface: %s\n", current.Name)
		fmt.Printf("Bytes Sent: %d\n", current.BytesSent-initial.BytesSent)
		fmt.Printf("Bytes Received: %d\n", current.BytesRecv-initial.BytesRecv)
		fmt.Printf("Packets Sent: %d\n", current.PacketsSent-initial.PacketsSent)
		fmt.Printf("Packets Received: %d\n", current.PacketsRecv-initial.PacketsRecv)
		fmt.Printf("Errors In: %d\n", current.Errin-initial.Errin)
		fmt.Printf("Errors Out: %d\n", current.Errout-initial.Errout)
		fmt.Printf("Drop In: %d\n", current.Dropin-initial.Dropin)
		fmt.Printf("Drop Out: %d\n", current.Dropout-initial.Dropout)
		fmt.Println("-----------------------------")
	}
}
