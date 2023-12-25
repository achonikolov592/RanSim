package main

import (
	"fmt"
	"helpers"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type location struct {
	path string
	name string
}

var whichTests []int
var testLocation = []location{location{"../ArchiveFiles/", "ArchiveFiles"}, location{"../Eicar/", "Eicar"}, location{"../EncryptDecryptDirRecursive/", "EncryptDecryptDirRecursive"}, location{"../EncryptDecryptDirRecursivePartially/", "EncryptDecryptDirRecursivePartially"}, location{"../SecureDeleteFiles/", "SecureDeleteFile"}, location{"../StartupFolderNewFile/", "startup"}}

var nameOfLogFile string

func prepareEveryTest() {
	err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		if path != "../" {
			info, err := os.Stat(path)
			if err != nil {
				return err
			} else if info.IsDir() && path != "../main" && path[:7] != "../.git" && !strings.Contains(path, "testfiles") {
				cmd := exec.Command("go", "build", ".")
				cmd.Dir = path
				err = cmd.Run()
				if err != nil {
					helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
					os.Exit(2)
				}
			}
		}
		return nil
	})

	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(1)
	}
}

func main() {
	nameOfLogFile := helpers.CreateLogFileIfItDoesNotExist("./", "main")

	prepareEveryTest()

	fmt.Println("These are the tests:\n1.ArchiveFiles\n2.Eicar File Creation\n3.EncryptDecryptDirRecursive\n4.EncryptDecryptDirRecursivePartially\n5.SecureDeleteFiles\nStartupFolderNewFile\nWhich one do you want: ")
	var i = 1
	fmt.Scanf("%d", &i)
	for i != 0 {
		whichTests = append(whichTests, i)
		fmt.Scanf("%d", &i)
	}

	for i = 0; i < len(whichTests); i++ {
		if whichTests[i] == 3 || whichTests[i] == 4 {
			var encr, decr string
			fmt.Println("Choose options for encrypt and decrypt:")
			fmt.Scanf("%s", &encr)
			fmt.Scanf("%s", &decr)
			cmd := exec.Command(testLocation[whichTests[i]-1].path+testLocation[whichTests[i]-1].name, encr, decr)
			cmd.Dir = testLocation[whichTests[i]-1].path
			err := cmd.Run()
			if err != nil {
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
			}
		} else {
			cmd := exec.Command(testLocation[whichTests[i]-1].path + testLocation[whichTests[i]-1].name)
			cmd.Dir = testLocation[whichTests[i]-1].path
			err := cmd.Run()
			if err != nil {
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
			}
		}
	}
}
