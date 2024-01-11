package obfuscate

import (
	"fmt"
	"helpers"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func PrepareEveryTestObfuscated(nameOfLogFile string) {
	err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		if path != "../" {
			_, err := os.Stat(path)
			if err != nil {
				return err
			} else if info.IsDir() && !strings.Contains(path, "main") && !strings.Contains(path, ".git") && !strings.Contains(path, "testfiles") {
				fmt.Println("Building: " + path)
				cmd := exec.Command("garble", "-literals", "build", ".")
				cmd.Dir = path
				err = cmd.Run()
				if err != nil {
					helpers.WriteLog(nameOfLogFile, err.Error(), 1)
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(1)
	}
}
