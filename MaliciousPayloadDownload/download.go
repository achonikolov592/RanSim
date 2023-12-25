package main

import (
	"helpers"
	"io"
	"net/http"
	"os"
)

var nameOfLogFile string

func main() {
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "eicar")

	helpers.WriteLog(nameOfLogFile, "Starting test: DownloadingMaliciousPayload")

	out, err := os.Create("/out1.exe")
	if err != nil {
		os.Exit(1)
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
	}
	defer out.Close()

	resp, err := http.Get("https://bazaar.abuse.ch/download/124e57afaa34fee7a6405531b05ceb4466e94f9e1858db976188a36de58156b3/")
	if err != nil {
		os.Exit(2)
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(3)
	}

	_, err = os.Open("./out1.exe")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(4)
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: DownloadingMaliciousPayload")
}

//
