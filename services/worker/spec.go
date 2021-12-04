package worker

import (
	"errors"
	"fmt"

	"github.com/quarterblue/beehive/services/worker/pb"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	ErrHost error = errors.New("machine spec: could not fetch host information")
	ErrDisk error = errors.New("machine spec: could not fetch disk information")
	ErrMem  error = errors.New("machine spec: could not fetch memory information")
	ErrCpu  error = errors.New("machine spec: could not fetch cpu information")
)

func ByteCountSI(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func retrieveSpec() (*pb.SpecResponse, error) {

	hInfo, err := host.Info()
	if err != nil {
		return nil, ErrHost
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, ErrDisk
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, ErrMem
	}

	cpu, err := cpu.Info()
	if err != nil {
		return nil, ErrCpu
	}

	cpuCount := uint64(len(cpu))

	return &pb.SpecResponse{
		CPUmodel:   cpu[0].ModelName,
		CPUmhz:     cpu[0].Mhz,
		CPUcore:    cpuCount,
		Memory:     v.Total,
		MemoryFree: v.Free,
		Disk:       diskStat.Total,
		DiskFree:   diskStat.Free,
		OS:         hInfo.OS,
		Platform:   hInfo.Platform,
		KernelArch: hInfo.KernelArch,
		Hostname:   hInfo.Hostname,
		Uptime:     hInfo.Uptime,
		BootTime:   hInfo.BootTime,
	}, nil
}
