package main

import (
	"helpers"
	"io"
	"os"
	"os/exec"
)

func main() {
	name := helpers.CreateLogFileIfItDoesNotExist("./", "startup")
	helpers.CreateLogFileIfItDoesNotExist("./", "EncryptionInfo")
	helpers.WriteLog(name, "Strating test : StartupFolderNewFile")
	helpers.CreateTestFiles("./", name)

	compileFile := exec.Command("go", "build", ".")
	compileFile.Dir = "./encr"

	err := compileFile.Run()
	if err != nil {
		helpers.WriteLog(name, "Error: "+err.Error())
		os.Exit(1)
	}

	src, err := os.Open("./encr/encr.exe")
	if err != nil {
		helpers.WriteLog(name, "Error: "+err.Error())
		os.Exit(2)
	}

	dest, err := os.Create("C:/Users/achon/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/a.exe")
	if err != nil {
		helpers.WriteLog(name, "Error: "+err.Error())
		os.Exit(3)
	}

	_, err = io.Copy(dest, src)
	if err != nil {
		helpers.WriteLog(name, "Error: "+err.Error())
		os.Exit(4)
	}

	helpers.WriteLog("./startup.log", "Ending test: StartupFolderNewFile")
}

//
