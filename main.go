package main

import (
	"fmt"
	"os"
	"os/exec"
)

func addToFile(){
	file, _ := os.OpenFile("./Registries/regs.go", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	bytesToAdd := []byte("a")
	
	file.Write(bytesToAdd)
}

func main() {
	// addToFile()
	
	// cmd1 := exec.Command("go", "build", "./")
	// cmd1.Dir = "./Registries"
	
	cmd2 := exec.Command("./Registries/registry")	
	cmd2.Stdout = os.Stdout
	

	/*if err := cmd1.Run(); err != nil {
		fmt.Println(err)
	}*/

	if err := cmd2.Run(); err != nil {
		fmt.Println(err)
	}
}