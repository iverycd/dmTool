package main

// This example requires go-cmd v1.2 or newer

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"

	"github.com/go-cmd/cmd"
)

func main() {
	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, "cmd", "/C", "dm_client\\dimp test/Gepoint@192.168.74.10 file=test.dmp table_exists_action=replace schemas=test")

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		// https://dave.cheney.net/2013/04/30/curious-channels
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				//fmt.Println(line)
				//将GBK编码的字符串转换为utf-8编码
				output, err := simplifiedchinese.GBK.NewDecoder().String(line)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(output)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-envCmd.Start()

	// Wait for goroutine to print everything
	<-doneChan
}
