package hardware

import (
	"fmt"
	"os/exec"
	// "path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type HardwareMetrics struct {
	CPUs     int
	TotalRam uint64
	FreeRAM  uint64
	Disks    []DiskInfo
}

type DiskInfo struct {
	Name  string
	Space uint64
	Free  uint64
	Mount string
}

func GetHardwareMetrics() (*HardwareMetrics, error) {
	metrics := &HardwareMetrics{}
	metrics.CPUs = runtime.NumCPU()
	// metrics.CPUs=0

	totalRam, err := GetTotalRAM()
	if err != nil {
		return nil, err
	}
	metrics.TotalRam = totalRam

	freeRam, err := GetFreeRAM()
	if err != nil {
		return nil, err
	}
	metrics.FreeRAM = freeRam

	disks, err := GetDiskInfo()
	if err != nil {
		return nil, err
	}
	metrics.Disks = disks

	return metrics, nil
}

func GetTotalRAM() (uint64, error) {
	var out []byte
	var err error

	if runtime.GOOS == "windows" {
		out, err = exec.Command("wmic", "ComputerSystem", "get", "TotalPhysicalMemory", "/Value").Output()
		// fmt.Printf(string(out))
	} else { // Unix
		out, err = exec.Command("free", "-b").Output()
	}

	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return 0, fmt.Errorf("unexpected error from terminal :(")
	}


	var total uint64
	if runtime.GOOS == "windows" {
		total,err=strconv.ParseUint(strings.Split(strings.TrimSpace(string(out)), "=")[1],10,64)
	} else { // Unix
		fields := strings.Fields(lines[1])
		if len(fields) < 2 {
			return 0, fmt.Errorf("Unexpected error from terminal :(")
		}
		total, err = strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return 0, err
		}
	}

	return total/1024/1024/1024, nil
}

func GetFreeRAM() (uint64, error) {
	var out []byte
	var err error

	if runtime.GOOS == "windows" {
		out, err = exec.Command("wmic", "OS", "get", "FreePhysicalMemory", "/Value").Output()
		// fmt.Println("hehehe",strings.Split(strings.TrimSpace(string(out)), "=")[1],"error is")
		

	} else { // Unix
		out, err = exec.Command("free", "-b").Output()
	}
	
	if err != nil {
		return 0, err
	}

	// lines := strings.Split(string(out), "\n")
	// if len(lines) < 2 {
	// 	// fmt.Println("lines are",lines)
	// 	return 0, fmt.Errorf("Unexpected error from terminal line 110 :(")

	// }

	var free uint64
	if runtime.GOOS == "windows" {
		free,err=strconv.ParseUint(strings.Split(strings.TrimSpace(string(out)), "=")[1],10,64)
	} else { // Unix
		lines := strings.Split(string(out), "\n")
		fields := strings.Fields(lines[1])
		if len(fields) < 4 {
			return 0, fmt.Errorf("Unexpected error from terminal :(",fields,err)
		}
		free, err = strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return 0, err
		}
	}

	return free/1024/1024, nil
}

func GetDiskInfo() ([]DiskInfo, error) {
		var out []byte
		var err error

		if runtime.GOOS == "windows" {
			out, err = exec.Command("wmic", "logicaldisk", "get", "DeviceID,Size,FreeSpace,VolumeName", "/Format:List").Output()
		} else { // Unix
			out, err = exec.Command("df", "-BM").Output()
		}
		// fmt.Println("disk info is : ",strings.Split(strings.TrimSpace(string(out)), ""));
		// fmt.Println(string(out));

		// for _, v := range strings.Split(strings.TrimSpace(string(out)), "="){
			// fmt.Println("element is:",v);
		// }
		if err != nil {
			return nil, err
		}
		var lines []string
		if runtime.GOOS == "windows" {
			lines = strings.Split(string(out), "\r\n")
		} else {
			lines = strings.Split(string(out), "\n")
		}
	
		disks := make([]DiskInfo, 0)
		var currentDisk DiskInfo
		for _, line := range lines {
			fields := strings.Split(line, "=")
			// fmt.Println(fields);
			// fmt.Println("iteration has",strings.TrimSpace(fields[0]))
			if len(fields) != 2 {
				continue
			}
			switch strings.TrimSpace(fields[0]) {
			case "DeviceID":
				if currentDisk.Name != "" {
					disks = append(disks, currentDisk)
				}
				currentDisk = DiskInfo{Name: strings.TrimSpace(fields[1])}
				// fmt.Println("appening",strings.TrimSpace(fields[1]))
			case "Size":
				sizeStr := strings.TrimSpace(strings.TrimSuffix(fields[1], "M"))
				size, err := strconv.ParseUint(sizeStr, 10, 64)
				if err != nil {
					// fmt.Println(err.Error())
					// return 0, nil
					currentDisk.Space = 0
					continue
				}
				currentDisk.Space = size/1024/1024/1024
			case "FreeSpace":
				freeStr := strings.TrimSpace(strings.TrimSuffix(fields[1], "M"))
				free, err := strconv.ParseUint(freeStr, 10, 64)
				if err != nil {
					currentDisk.Free = 0
					continue
				}
				
				currentDisk.Free = free/1024/1024/1024
			case "VolumeName":
				currentDisk.Mount = strings.TrimSpace(fields[1])
			}
		}
		// Append the last disk
		if currentDisk.Name != "" {
			disks = append(disks, currentDisk)
		}
		return disks, nil
	
		// return disks, nil
}
