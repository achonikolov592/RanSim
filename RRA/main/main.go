package main

import (
<<<<<<< HEAD
	//obfuscate "RRA/Obfuscation"
	"fmt"
	"helpers"
	"os/exec"
=======
	"fmt"
	"helpers"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
)

type location struct {
	path string
	name string
}

<<<<<<< HEAD
var testLocation = []location{location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/ArchiveFiles/", "ArchiveFiles.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/Eicar/", "Eic.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/", "EncryptDecryptDirRecursive.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/", "EncryptDecryptDirRecursivePartially.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/GetSysInfo/", "GetSysInfo.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/SecureDeleteFiles/", "SecureDeleteFiles.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/", "Startup.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/", "MaliciousPayloadDownload.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/ServiceCreation/", "ServiceCreation.exe"}, location{"C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/PrivEsc/", "AccsTok.exe"}}
var testRightResult = []string{"nil", "exit status 1", "nil", "nil", "nil", "nil", "exit status 1", "exit status 1", "nil", "exit status 1"}

func main() {
	whichTests := make([]int, 0)

	nameOfLogFile := helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/main/", "main")

	//obfuscate.PrepareEveryTestObfuscated(nameOfLogFile)

	fmt.Println("These are the tests:\n-1.All\n1.ArchiveFiles\n2.Eicar File Creation\n3.EncryptDecryptDirRecursive\n4.EncryptDecryptDirRecursivePartially\n5.GetSysInfo\n6.SecureDeleteFiles\n7.StartupFolderNewFile\n8.MaliciousPayloadDownload\n9.ServiceCreation\n10.PrivilegeEscalation\nWhich one do you want: ")
	var i int
	fmt.Scanln(&i)
	for i != 0 {
		whichTests = append(whichTests, i)
		fmt.Scanln(&i)
	}

	var (
		correctTests   = 0
		incorrectTests = 0
	)
	//print tests results on every test
	if whichTests[0] == -1 {
		for i = 0; i < len(testLocation); i++ {
			fmt.Println("Starting test: " + testLocation[i].name)
			if i == 2 || i == 3 {
				var encr, decr string
				fmt.Println("Choose options for encrypt and decrypt:")
				fmt.Scanln(&encr)
				fmt.Scanln(&decr)
				cmd := exec.Command(testLocation[i].path+testLocation[i].name, encr, decr)
				cmd.Dir = testLocation[i].path
				err := cmd.Run()

				if err == nil {
					correctTests++
				} else {
					helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
					incorrectTests++
				}
			} else if i == 8 {
				var option string
				fmt.Println("Choose options to install, run or both:")
				fmt.Scanln(&option)
				if option != "both" {
					cmd := exec.Command(testLocation[i].path+testLocation[i].name, option)
					cmd.Dir = testLocation[i].path
					err := cmd.Run()

					if err == nil {
						if testRightResult[i] == "nil" {
							correctTests++
						} else {
							incorrectTests++
						}
					} else {
						if err.Error() == testRightResult[i] {
							correctTests++
						} else {
							helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
							incorrectTests++
						}
					}
				} else {
					cmd1 := exec.Command(testLocation[i].path+testLocation[i].name, "install")
					cmd1.Dir = testLocation[i].path
					cmd2 := exec.Command(testLocation[i].path+testLocation[i].name, "run")
					cmd2.Dir = testLocation[i].path
					err1 := cmd1.Run()
					err2 := cmd2.Run()

					if err1 == nil && err2 == nil {
						correctTests++
					} else {
						helpers.WriteLog(nameOfLogFile, err1.Error()+" from "+testLocation[i].name, 1)
						helpers.WriteLog(nameOfLogFile, err2.Error()+" from "+testLocation[i].name, 1)
						incorrectTests++
					}
				}
			} else if i == 9 {
				var pid string
				fmt.Println("Choose which pid to duplicate:")
				fmt.Scanln(&pid)
				cmd := exec.Command(testLocation[i].path+testLocation[i].name, pid)
				cmd.Dir = testLocation[i].path
				err := cmd.Run()

				if err.Error() == testRightResult[i] {
					correctTests++
				} else {
					helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
					incorrectTests++
				}
			} else {
				cmd := exec.Command(testLocation[i].path + testLocation[i].name)
				cmd.Dir = testLocation[i].path
				err := cmd.Run()

				if err == nil {
					if testRightResult[i] == "nil" {
						correctTests++
					} else {
						incorrectTests++
					}
				} else {
					if err.Error() == testRightResult[i] {
						correctTests++
					} else {
						helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
						incorrectTests++
					}
				}
			}
		}
		fmt.Printf("Out of %d tests, %d are/is correct, %d are/is incorrect", len(testLocation), correctTests, incorrectTests)
	} else {
		for i = 0; i < len(whichTests); i++ {
			fmt.Println("Starting Test: " + testLocation[i].name)
			if whichTests[i] == 3 || whichTests[i] == 4 {
				var encr, decr string
				fmt.Println("Choose options for encrypt and decrypt:")
				fmt.Scanln(&encr)
				fmt.Scanln(&decr)
				cmd := exec.Command(testLocation[whichTests[i]-1].path+testLocation[whichTests[i]-1].name, encr, decr)
				cmd.Dir = testLocation[whichTests[i]-1].path
				err := cmd.Run()

				if err == nil {
					correctTests++
				} else {
					helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
					incorrectTests++
				}
			} else if whichTests[i] == 9 {
				var option string
				fmt.Println("Choose options to install, run or both:")
				fmt.Scanln(&option)
				if option != "both" {
					cmd := exec.Command(testLocation[whichTests[i]-1].path+testLocation[whichTests[i]-1].name, option)
					cmd.Dir = testLocation[whichTests[i]-1].path
					err := cmd.Run()

					if err == nil {
						if testRightResult[i] == "nil" {
							correctTests++
						} else {
							incorrectTests++
						}
					} else {
						if err.Error() == testRightResult[i] {
							correctTests++
						} else {
							helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
							incorrectTests++
						}
					}
				} else {
					cmd1 := exec.Command(testLocation[whichTests[i]-1].path+testLocation[whichTests[i]-1].name, "run")
					cmd1.Dir = testLocation[whichTests[i]-1].path
					cmd2 := exec.Command(testLocation[whichTests[i]-1].path+testLocation[whichTests[i]-1].name, "run")
					cmd2.Dir = testLocation[whichTests[i]-1].path
					err1 := cmd1.Run()
					err2 := cmd2.Run()

					if err1 == nil && err2 == nil {
						correctTests++
					} else {
						helpers.WriteLog(nameOfLogFile, err1.Error()+" from "+testLocation[i].name, 1)
						helpers.WriteLog(nameOfLogFile, err2.Error()+" from "+testLocation[i].name, 1)
						incorrectTests++
					}
				}
			} else if whichTests[i] == 10 {
				var pid string
				fmt.Println("Choose which pid to duplicate:")
				fmt.Scanln(&pid)
				cmd := exec.Command(testLocation[whichTests[i]-1].path+testLocation[whichTests[i]-1].name, pid)
				cmd.Dir = testLocation[whichTests[i]-1].path
				err := cmd.Run()

				if err == nil {
					if testRightResult[i] == "nil" {
						correctTests++
					} else {
						incorrectTests++
					}
				} else {
					if err.Error() == testRightResult[i] {
						correctTests++
					} else {
						helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
						incorrectTests++
					}
				}
			} else {
				cmd := exec.Command(testLocation[whichTests[i]-1].path + testLocation[whichTests[i]-1].name)
				cmd.Dir = testLocation[whichTests[i]-1].path
				err := cmd.Run()

				if err == nil {
					if testRightResult[i] == "nil" {
						correctTests++
					} else {
						incorrectTests++
					}
				} else {
					if err.Error() == testRightResult[i] {
						correctTests++
					} else {
						helpers.WriteLog(nameOfLogFile, err.Error()+" from "+testLocation[i].name, 1)
						incorrectTests++
					}
				}
			}
		}
		fmt.Printf("Out of %d tests, %d are/is correct, %d are/is incorrect", len(whichTests), correctTests, incorrectTests)
	}

=======
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
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
}
