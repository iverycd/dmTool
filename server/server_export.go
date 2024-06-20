////go:build exp

package server

import (
	"dmTool/global"
	"fmt"
	"github.com/go-cmd/cmd"
	"golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

// ExpImp 导出
func ExpImp2() {
	fileName := fmt.Sprintf("%s_%s", global.Config.Server.User, time.Now().Format("2006_01_02_15_04_05"))
	// 获取服务端连接方式
	dmCmd := fmt.Sprintf("dm_client\\dexp %s/%s@%s:%s file=%s.dmp log=%s.log schemas=%s",
		global.Config.Server.User,
		global.Config.Server.Password,
		global.Config.Server.Host,
		global.Config.Server.Port,
		fileName,
		fileName,
		global.Config.Server.User,
	)
	fmt.Println(dmCmd)
	// Start a long-running process, capture stdout and stderr
	// 导出
	findCmd := cmd.NewCmd("cmd", "/C", dmCmd)
	// 导入
	//findCmd := cmd.NewCmd("cmd", "/C", "dm_client\\dimp test/Gepoint@192.168.74.10 file=dexp.dmp table_exists_action=replace schemas=test")
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
	fmt.Println("\nExport  Success in your current path ->", fileName+".dmp")
}
