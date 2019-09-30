package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

func main() {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	fmt.Println(virtualMemory)

	cpuInfo, err := cpu.Info()
	if err != nil {
		panic(err)
	}
	cpuCounts, err := cpu.Counts(true)
	if err != nil {
		panic(err)
	}
	cpuTimes, err := cpu.Times(true)
	if err != nil {
		panic(err)
	}
	fmt.Println(cpuInfo, cpuCounts, cpuTimes)

	loadAvg, err := load.Avg()
	if err != nil {
		panic(err)
	}
	loadMisc, err := load.Misc()
	if err != nil {
		panic(err)
	}
	fmt.Println(loadAvg, loadMisc)

	diskPartitions, err := disk.Partitions(true)
	if err != nil {
		panic(err)
	}
	fmt.Println(diskPartitions)

	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}
	processPid, err := process.Pids()
	if err != nil {
		panic(err)
	}
	fmt.Println(processes, processPid)

	hostInfo, err := host.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostInfo)

	dockerIDList, err := docker.GetDockerIDList()
	if err != nil {
		panic(err)
	}
	dockerStat, err := docker.GetDockerStat()
	if err != nil {
		panic(err)
	}
	fmt.Println(dockerIDList, dockerStat)

	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	fmt.Println(netInterfaces)
}
