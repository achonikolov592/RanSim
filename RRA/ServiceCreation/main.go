package main

import (
	"RRA/EncryptDecryptDirRecursive/encrypt"
	"helpers"
	"os"
	"time"

	"github.com/kardianos/service"
)

var nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/onedrive/desktop/diplomna1/rra/ServiceCreation/", "ServiceCreation")
var nameOfEncryptionInfoFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/onedrive/desktop/diplomna1/rra/ServiceCreation/", "EncryptionInfo")

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	for {
		helpers.WriteLog(nameOfLogFile, "Starting Encryption", 2)
		helpers.CreateTestFiles("C:/Users/achon/onedrive/desktop/diplomna1/rra/ServiceCreation/", nameOfLogFile)
		encrypt.EncryptDir("C:/Users/achon/onedrive/desktop/diplomna1/rra/ServiceCreation/testfiles/", nameOfLogFile, nameOfEncryptionInfoFile)
		helpers.WriteLog(nameOfLogFile, "Ending Encryption", 2)
		time.Sleep(10 * time.Second)
	}
}
func (p *program) Stop(s service.Service) error {
	return s.Stop()
}

func main() {
	if len(os.Args) != 2 {
		helpers.WriteLog(nameOfLogFile, "Invalid number of arguments!", 1)
		os.Exit(1)
	}

	helpers.WriteLog(nameOfLogFile, "Starting test: ServiceCreation", 2)
	options := make(service.KeyValue)
	options[service.StartType] = service.ServiceStartManual

	svcConfig := &service.Config{
		Name:        "EncrFile",
		DisplayName: "EncryptDir",
		Option:      options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)

	if err != nil {
		os.Exit(1)
	}

	if os.Args[1] == "install" {
		helpers.WriteLog(nameOfLogFile, "Installing service", 2)
		err := s.Install()
		if err != nil {
			helpers.WriteLog(nameOfLogFile, err.Error(), 1)
			os.Exit(2)
		}
		helpers.WriteLog(nameOfLogFile, "Installed service", 2)
	} else if os.Args[1] == "uninstall" {
		helpers.WriteLog(nameOfLogFile, "Uninstalling service", 2)
		err := s.Uninstall()
		if err != nil {
			helpers.WriteLog(nameOfLogFile, err.Error(), 1)
			os.Exit(3)
		}
		helpers.WriteLog(nameOfLogFile, "Uninstalled service", 2)
	} else if os.Args[1] == "run" {
		helpers.WriteLog(nameOfLogFile, "Running service", 2)
		helpers.WriteLog(nameOfLogFile, s.Run().Error(), 1)
		helpers.WriteLog(nameOfLogFile, "Stopped running service", 2)
	}

	os.Exit(0)
}
