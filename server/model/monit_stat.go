package model

import "time"

type FSInfo struct {
	MountPoint string `json:"MountPoint"`
	Used       uint64 `json:"Used"`
	Free       uint64 `json:"Free"`
}

type NetIntfInfo struct {
	IPv4 string `json:"IPv4"`
	IPv6 string `json:"IPv6"`
	Rx   uint64 `json:"Rx"`
	Tx   uint64 `json:"Tx"`
}

type cpuRaw struct {
	User    uint64 `json:"User"`    // time spent in user mode
	Nice    uint64 `json:"Nice"`    // time spent in user mode with low priority (nice)
	System  uint64 `json:"System"`  // time spent in system mode
	Idle    uint64 `json:"Idle"`    // time spent in the idle task
	Iowait  uint64 `json:"Iowait"`  // time spent waiting for I/O to complete (since Linux 2.5.41)
	Irq     uint64 `json:"Irq"`     // time spent servicing  interrupts  (since  2.6.0-test4)
	SoftIrq uint64 `json:"SoftIrq"` // time spent servicing softirqs (since 2.6.0-test4)
	Steal   uint64 `json:"Steal"`   // time spent in other OSes when running in a virtualized environment
	Guest   uint64 `json:"Guest"`   // time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.
	Total   uint64 `json:"Total"`   // total of all time fields
}

type CPUInfo struct {
	User    float32 `json:"User"`
	Nice    float32 `json:"Nice"`
	System  float32 `json:"System"`
	Idle    float32 `json:"Idle"`
	Iowait  float32 `json:"Iowait"`
	Irq     float32 `json:"Irq"`
	SoftIrq float32 `json:"SoftIrq"`
	Steal   float32 `json:"Steal"`
	Guest   float32 `json:"Guest"`
}

type MonitorStat struct {
	Uptime       time.Duration          `json:"Uptime"`
	Hostname     string                 `json:"Hostname"`
	Load1        string                 `json:"Load1"`
	Load5        string                 `json:"Load5"`
	Load10       string                 `json:"Load10"`
	RunningProcs string                 `json:"RunningProcs"`
	TotalProcs   string                 `json:"TotalProcs"`
	MemTotal     uint64                 `json:"MemTotal"`
	MemFree      uint64                 `json:"MemFree"`
	MemBuffers   uint64                 `json:"MemBuffers"`
	MemCached    uint64                 `json:"MemCached"`
	SwapTotal    uint64                 `json:"SwapTotal"`
	SwapFree     uint64                 `json:"SwapFree"`
	FSInfos      []FSInfo               `json:"FSInfos"`
	NetIntf      map[string]NetIntfInfo `json:"NetIntf"`
	CPU          CPUInfo                `json:"CPU"` // or []CPUInfo to get all the cpu-core's stats?
}
