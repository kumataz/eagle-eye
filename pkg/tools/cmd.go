package tools

import (
	"os/exec"
	"fmt"
	"io/ioutil"
)


// Get linux-command and return work result
func ParseCMD(intdata string) []uint8{
	input := intdata
	cmd := exec.Command("/bin/bash", "-c", input)
	// Create get command output pipeline
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
	}

	// execute cmd
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
	}

	// read all ouput
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
	}

	// output
	return bytes
	// fmt.Printf("%s\n", bytes)
}

