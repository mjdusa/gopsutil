// +build linux

package gopsutil

import (
	"errors"
	"strconv"
	"strings"
)

func Cpu_times() ([]CPU_TimesStat, error) {
	ret := make([]CPU_TimesStat, 0)

	filename := "/proc/stat"
	lines, _ := ReadLines(filename)
	for _, line := range lines {
		ct, err := parseStatLine(line)
		if err != nil {
			continue
		}
		ret = append(ret, ct)

	}
	return ret, nil
}

func parseStatLine(line string) (CPU_TimesStat, error) {
	fields := strings.Fields(line)

	if strings.HasPrefix(fields[0], "cpu") == false {
		//		return CPU_TimesStat{}, e
		return CPU_TimesStat{}, errors.New("not contain cpu")
	}

	cpu := fields[0]
	if cpu == "cpu" {
		cpu = "cpu-total"
	}
	user, _ := strconv.ParseFloat(fields[1], 32)
	nice, _ := strconv.ParseFloat(fields[2], 32)
	system, _ := strconv.ParseFloat(fields[3], 32)
	idle, _ := strconv.ParseFloat(fields[4], 32)
	iowait, _ := strconv.ParseFloat(fields[5], 32)
	irq, _ := strconv.ParseFloat(fields[6], 32)
	softirq, _ := strconv.ParseFloat(fields[7], 32)
	stolen, _ := strconv.ParseFloat(fields[8], 32)
	ct := CPU_TimesStat{
		Cpu:     cpu,
		User:    float32(user),
		Nice:    float32(nice),
		System:  float32(system),
		Idle:    float32(idle),
		Iowait:  float32(iowait),
		Irq:     float32(irq),
		Softirq: float32(softirq),
		Stolen:  float32(stolen),
	}
	if len(fields) > 9 { // Linux >= 2.6.11
		steal, _ := strconv.ParseFloat(fields[9], 32)
		ct.Steal = float32(steal)
	}
	if len(fields) > 10 { // Linux >= 2.6.24
		guest, _ := strconv.ParseFloat(fields[10], 32)
		ct.Guest = float32(guest)
	}
	if len(fields) > 11 { // Linux >= 3.2.0
		guest_nice, _ := strconv.ParseFloat(fields[11], 32)
		ct.Guest_nice = float32(guest_nice)
	}

	return ct, nil
}