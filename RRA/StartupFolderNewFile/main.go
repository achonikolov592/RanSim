package main

import (
	"helpers"
	"io"
	"os"
	"os/exec"
)

func main() {
<<<<<<< HEAD
	name := helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/", "startup")
	helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/", "EncryptionInfo")
	helpers.WriteLog(name, "Strating test : StartupFolderNewFile", 2)
	
	compileFile := exec.Command("go", "build", ".")
	compileFile.Dir = "C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/encr"

	err := compileFile.Run()
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(2)
	}
	src, err := os.Open("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/encr/enc.exe")
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(3)
	}
	dest, err := os.Create("C:/Users/achon/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/a.exe")
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(4)
	}
	_, err = io.Copy(dest, src)
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(5)
	}

	_, err = os.Open("C:/Users/achon/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/a.exe")
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(1)
	}

	helpers.WriteLog(name, "Ending test: StartupFolderNewFile", 2)
	os.Exit(0)
}

//aaa
=======
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
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
