package capture

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"net"
	"time"
)


const (
	NUM_KB = 1000.0000
	NUM_MIB = 1000000.0000
	NUM_GB = 1000000000.0000
)

func HanderUnit(num uint64, numtype float64,typename string) (newnum string) {

	f :=fmt.Sprintf("%.2f", float64(num)/numtype)
	return f + typename
}


func GetSystemInfo() map[string]string {
	systemdata := make(map[string]string)

	// cpu
	cpuPercent,_ := cpu.Percent(time.Second,true)
	// cpuNumber,_ := cpu.Counts(true)
	// systemdata["cpu_cores"] = fmt.Sprint(cpuNumber)
	systemdata["cpu_usage"] = fmt.Sprintf("%.2f%%", cpuPercent[0])
	// mem
	v, _ := mem.VirtualMemory()
	// total := HanderUnit(v.Total,NUM_GB,"G")
	// used := HanderUnit(v.Used,NUM_GB,"G")
	// free := HanderUnit(v.Free,NUM_GB,"G")
	userPercent := fmt.Sprintf("%.2f",v.UsedPercent)
	// systemdata["mem_total"] = total
	// systemdata["mem_used"] = used
	// systemdata["mem_free"] = free
	systemdata["mem_usage"] = userPercent + "%"
	// host
	localIP, _ := GetLocalIP()
	systemdata["ip_address"] = localIP
	systemdata["machine_id"] = localIP[10:]
	// hInfo, _ := host.Info()
	// systemdata["pc_hostname"] = hInfo.Hostname
	// systemdata["pc_os"] = hInfo.OS
	// systemdata["pc_platform"] = hInfo.Platform + "-" + hInfo.PlatformVersion

	return systemdata
}


// cpu info
func GetCpuInfo() map[string]string {
	cpuPercent,_ := cpu.Percent(time.Second,true)
	cpuNumber,_ := cpu.Counts(true)
	cpudata := make(map[string]string)
	cpudata["cores"] = fmt.Sprint(cpuNumber)
	cpudata["Usage"] = fmt.Sprintf("%.2f%%", cpuPercent[0])
	return cpudata
}

// memory info
func GetMemInfo() (map[string]string) {	
	memdata := make(map[string]string)
	v, _ := mem.VirtualMemory()

	total := HanderUnit(v.Total,NUM_GB,"G")
	used := HanderUnit(v.Used,NUM_GB,"G")
	free := HanderUnit(v.Free,NUM_GB,"G")
	userPercent := fmt.Sprintf("%.2f",v.UsedPercent)

	memdata["Total"] = total
	memdata["Used"] = used
	memdata["Free"] = free
	memdata["Usage"] = userPercent + "%"

	return memdata
}

// ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// host info
func GetHostInfo() map[string]string {
	hostdata := make(map[string]string)

	localIP, _ := GetLocalIP()
	hostdata["ip"] = localIP

	hInfo, _ := host.Info()
	hostdata["hostname"] = hInfo.Hostname
	hostdata["os"] = hInfo.OS
	hostdata["platform"] = hInfo.Platform + "-" + hInfo.PlatformVersion

	return hostdata
}


// disk
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}



