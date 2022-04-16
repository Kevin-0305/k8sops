package sshStat

import (
	"bufio"
	"encoding/json"
	"gin-vue-admin/global"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type FSInfo struct {
	MountPoint string `json:"mountPoint"`
	Used       uint64 `json:"used"`
	Free       uint64 `json:"free"`
}

type NetIntfInfo struct {
	IPv4 string `json:"ipv4"`
	IPv6 string `json:"ipv6"`
	Rx   uint64 `json:"rx"`
	Tx   uint64 `json:"tx"`
}

type cpuRaw struct {
	User    uint64 `json:"user"`    // time spent in user mode
	Nice    uint64 `json:"nice"`    // time spent in user mode with low priority (nice)
	System  uint64 `json:"system"`  // time spent in system mode
	Idle    uint64 `json:"idle"`    // time spent in the idle task
	Iowait  uint64 `json:"iowait"`  // time spent waiting for I/O to complete (since Linux 2.5.41)
	Irq     uint64 `json:"irq"`     // time spent servicing  interrupts  (since  2.6.0-test4)
	SoftIrq uint64 `json:"softIrq"` // time spent servicing softirqs (since 2.6.0-test4)
	Steal   uint64 `json:"steal"`   // time spent in other OSes when running in a virtualized environment
	Guest   uint64 `json:"guest"`   // time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.
	Total   uint64 `json:"total"`   // total of all time fields
}

type CPUInfo struct {
	User    float32 `json:"user"`
	Nice    float32 `json:"nice"`
	System  float32 `json:"system"`
	Idle    float32 `json:"idle"`
	Iowait  float32 `json:"iowait"`
	Irq     float32 `json:"irq"`
	SoftIrq float32 `json:"softIrq"`
	Steal   float32 `json:"steal"`
	Guest   float32 `json:"guest"`
}

type Stats struct {
	Uptime       time.Duration          `json:"uptime"`
	Hostname     string                 `json:"hostname"`
	Load1        string                 `json:"load1"`
	Load5        string                 `json:"load5"`
	Load10       string                 `json:"load10"`
	RunningProcs string                 `json:"runningProcs"`
	TotalProcs   string                 `json:"totalProcs"`
	MemTotal     uint64                 `json:"memTotal"`
	MemFree      uint64                 `json:"memFree"`
	MemBuffers   uint64                 `json:"memBuffers"`
	MemCached    uint64                 `json:"memCached"`
	SwapTotal    uint64                 `json:"swapTotal"`
	SwapFree     uint64                 `json:"swapFree"`
	FSInfos      []FSInfo               `json:"fsInfos"`
	NetIntf      map[string]NetIntfInfo `json:"netIntf"`
	CPU          CPUInfo                `json:"cpu"` // or []CPUInfo to get all the cpu-core's stats?
}

func GetStats(addr string, user string, password string, port int, hostStatsMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	var stats Stats
	client, err := ssh.Dial("tcp", addr+":"+strconv.Itoa(port), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
	}
	getAllStats(client, &stats)
	statsJson, err := json.Marshal(stats)
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
	}
	hostStatsMap.Store(addr, string(statsJson))
}

func getAllStats(client *ssh.Client, stats *Stats) {
	getUptime(client, stats)
	getHostname(client, stats)
	getLoad(client, stats)
	getMemInfo(client, stats)
	getFSInfo(client, stats)
	getInterfaces(client, stats)
	getInterfaceInfo(client, stats)
	getCPU(client, stats)
}

func getUptime(client *ssh.Client, stats *Stats) (err error) {

	uptime, err := runCommand(client, "/bin/cat /proc/uptime")
	if err != nil {
		return
	}

	parts := strings.Fields(uptime)
	if len(parts) == 2 {
		var upsecs float64
		upsecs, err = strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return
		}
		stats.Uptime = time.Duration(upsecs * 1e9)
	}

	return
}

func getHostname(client *ssh.Client, stats *Stats) (err error) {
	hostname, err := runCommand(client, "/bin/hostname -f")
	if err != nil {
		return
	}

	stats.Hostname = strings.TrimSpace(hostname)
	return
}

func getLoad(client *ssh.Client, stats *Stats) (err error) {
	line, err := runCommand(client, "/bin/cat /proc/loadavg")
	if err != nil {
		return
	}

	parts := strings.Fields(line)
	if len(parts) == 5 {
		stats.Load1 = parts[0]
		stats.Load5 = parts[1]
		stats.Load10 = parts[2]
		if i := strings.Index(parts[3], "/"); i != -1 {
			stats.RunningProcs = parts[3][0:i]
			if i+1 < len(parts[3]) {
				stats.TotalProcs = parts[3][i+1:]
			}
		}
	}

	return
}

func getMemInfo(client *ssh.Client, stats *Stats) (err error) {
	lines, err := runCommand(client, "/bin/cat /proc/meminfo")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 3 {
			val, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				continue
			}
			val *= 1024
			switch parts[0] {
			case "MemTotal:":
				stats.MemTotal = val
			case "MemFree:":
				stats.MemFree = val
			case "Buffers:":
				stats.MemBuffers = val
			case "Cached:":
				stats.MemCached = val
			case "SwapTotal:":
				stats.SwapTotal = val
			case "SwapFree:":
				stats.SwapFree = val
			}
		}
	}

	return
}

func getFSInfo(client *ssh.Client, stats *Stats) (err error) {
	lines, err := runCommand(client, "/bin/df -B1")
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	flag := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		n := len(parts)
		dev := n > 0 && strings.Index(parts[0], "/dev/") == 0
		if n == 1 && dev {
			flag = 1
		} else if (n == 5 && flag == 1) || (n == 6 && dev) {
			i := flag
			flag = 0
			used, err := strconv.ParseUint(parts[2-i], 10, 64)
			if err != nil {
				continue
			}
			free, err := strconv.ParseUint(parts[3-i], 10, 64)
			if err != nil {
				continue
			}
			stats.FSInfos = append(stats.FSInfos, FSInfo{
				parts[5-i], used, free,
			})
		}
	}

	return
}

func getInterfaces(client *ssh.Client, stats *Stats) (err error) {
	var lines string
	lines, err = runCommand(client, "/bin/ip -o addr")
	if err != nil {
		// try /sbin/ip
		lines, err = runCommand(client, "/sbin/ip -o addr")
		if err != nil {
			return
		}
	}

	if stats.NetIntf == nil {
		stats.NetIntf = make(map[string]NetIntfInfo)
	}

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 4 && (parts[2] == "inet" || parts[2] == "inet6") {
			ipv4 := parts[2] == "inet"
			intfname := parts[1]
			if info, ok := stats.NetIntf[intfname]; ok {
				if ipv4 {
					info.IPv4 = parts[3]
				} else {
					info.IPv6 = parts[3]
				}
				stats.NetIntf[intfname] = info
			} else {
				info := NetIntfInfo{}
				if ipv4 {
					info.IPv4 = parts[3]
				} else {
					info.IPv6 = parts[3]
				}
				stats.NetIntf[intfname] = info
			}
		}
	}

	return
}

func getInterfaceInfo(client *ssh.Client, stats *Stats) (err error) {
	lines, err := runCommand(client, "/bin/cat /proc/net/dev")
	if err != nil {
		return
	}

	if stats.NetIntf == nil {
		return
	} // should have been here already

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 17 {
			intf := strings.TrimSpace(parts[0])
			intf = strings.TrimSuffix(intf, ":")
			if info, ok := stats.NetIntf[intf]; ok {
				rx, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					continue
				}
				tx, err := strconv.ParseUint(parts[9], 10, 64)
				if err != nil {
					continue
				}
				info.Rx = rx
				info.Tx = tx
				stats.NetIntf[intf] = info
			}
		}
	}

	return
}

func parseCPUFields(fields []string, stat *cpuRaw) {
	numFields := len(fields)
	for i := 1; i < numFields; i++ {
		val, err := strconv.ParseUint(fields[i], 10, 64)
		if err != nil {
			continue
		}

		stat.Total += val
		switch i {
		case 1:
			stat.User = val
		case 2:
			stat.Nice = val
		case 3:
			stat.System = val
		case 4:
			stat.Idle = val
		case 5:
			stat.Iowait = val
		case 6:
			stat.Irq = val
		case 7:
			stat.SoftIrq = val
		case 8:
			stat.Steal = val
		case 9:
			stat.Guest = val
		}
	}
}

// the CPU stats that were fetched last time round
var preCPU cpuRaw

func getCPU(client *ssh.Client, stats *Stats) (err error) {
	lines, err := runCommand(client, "/bin/cat /proc/stat")
	if err != nil {
		return
	}

	var (
		nowCPU cpuRaw
		total  float32
	)

	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) > 0 && fields[0] == "cpu" { // changing here if want to get every cpu-core's stats
			parseCPUFields(fields, &nowCPU)
			break
		}
	}
	if preCPU.Total == 0 { // having no pre raw cpu data
		goto END
	}

	total = float32(nowCPU.Total - preCPU.Total)
	stats.CPU.User = float32(nowCPU.User-preCPU.User) / total * 100
	stats.CPU.Nice = float32(nowCPU.Nice-preCPU.Nice) / total * 100
	stats.CPU.System = float32(nowCPU.System-preCPU.System) / total * 100
	stats.CPU.Idle = float32(nowCPU.Idle-preCPU.Idle) / total * 100
	stats.CPU.Iowait = float32(nowCPU.Iowait-preCPU.Iowait) / total * 100
	stats.CPU.Irq = float32(nowCPU.Irq-preCPU.Irq) / total * 100
	stats.CPU.SoftIrq = float32(nowCPU.SoftIrq-preCPU.SoftIrq) / total * 100
	stats.CPU.Guest = float32(nowCPU.Guest-preCPU.Guest) / total * 100
END:
	preCPU = nowCPU
	return
}
