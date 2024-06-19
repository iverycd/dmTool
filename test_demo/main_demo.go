package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

// 第三方包调用cmd，能实时输出
func main() {
	// Start a long-running process, capture stdout and stderr
	// 导出
	//findCmd := cmd.NewCmd("cmd", "/C", "dm_client\\dexp test/Gepoint@192.168.74.10 schemas=test")
	// 导入
	findCmd := cmd.NewCmd("cmd", "/C", "dm_client\\dimp test/Gepoint@192.168.74.10 file=dexp.dmp table_exists_action=replace schemas=test")
	//findCmd := cmd.NewCmd("cmd", "/C", "dir")
	statusChan := findCmd.Start() // non-blocking

	ticker := time.NewTicker(10 * time.Millisecond)

	// Print last line of stdout every 2s
	go func() {
		for range ticker.C {
			status := findCmd.Status()
			n := len(status.Stdout)
			if n > 0 {
				//将GBK编码的字符串转换为utf-8编码
				output, err := simplifiedchinese.GBK.NewDecoder().String(status.Stdout[n-1])
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(output)
			}
		}
	}()

	// Stop command after 1 hour
	go func() {
		<-time.After(1 * time.Hour)
		findCmd.Stop()
	}()

	// Check if command is done
	select {
	case finalStatus := <-statusChan:
		// done
		fmt.Println(finalStatus)
	default:
		// no, still running
	}

	// Block waiting for command to exit, be stopped, or be killed
	finalStatus := <-statusChan
	n := len(finalStatus.Stdout)
	if n > 0 {
		//将GBK编码的字符串转换为utf-8编码
		output, err := simplifiedchinese.GBK.NewDecoder().String(finalStatus.Stdout[n-1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}
}

// go原生cmd调用方法
//func main() {
//	// 创建一个Cmd结构体
//	cmd := exec.Command("cmd", "/C", "dm_client\\dexp test/Gepoint@192.168.74.10 schemas=test")
//	// // 不需要cmd.Run()
//	out, err := cmd.Output()
//	//将GBK编码的字符串转换为utf-8编码
//	output, err := simplifiedchinese.GBK.NewDecoder().Bytes(out)
//	if err != nil {
//		fmt.Println("执行命令出错: ", err)
//		return
//	} else {
//		fmt.Println("获取命令执行结果: ", string(output))
//	}
//}
