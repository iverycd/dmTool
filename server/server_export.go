//go:build exp

package server

import (
	"dmTool/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-cmd/cmd"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"strings"
	"time"
)

// ExpImp 导出
func ExpImp() {
	var isConfirm string
	colorStr := color.New()
	colorStr.Add(color.FgHiRed)
	colorStr.Printf("目标主机-> %s \n目标数据库连接账号-> %s "+
		"\n请确认信息是否正确，输入\"YES\"继续,或者\"NO\"退出此程序\n", global.Config.Server.Host,
		global.Config.Server.User)
	// 只有在终端输入YES之后才会导出
	fmt.Scanln(&isConfirm)
	if strings.ToUpper(isConfirm) != "YES" {
		return
	}
	// 导出的文件名
	fileName := fmt.Sprintf("%s_%s", global.Config.Server.User, time.Now().Format("2006_01_02_15_04_05"))
	// 拼接导出命令
	dmCmd := fmt.Sprintf("dm_client\\dexp %s/\"%s\"@%s:%s file=%s.dmp log=%s.log schemas=%s",
		global.Config.Server.User,
		global.Config.Server.Password,
		global.Config.Server.Host,
		global.Config.Server.Port,
		fileName,
		fileName,
		global.Config.Server.User,
	)
	fmt.Println(dmCmd)

	// Disable output buffering, enable streaming 设置cmd输出的方式
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options 设置cmd选项
	envCmd := cmd.NewCmdOptions(cmdOptions, "cmd", "/C", dmCmd)

	// 执行导出

	// Print STDOUT and STDERR lines streaming from Cmd 实时输出命令行程序标准输出
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
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

	// Run and wait for Cmd to return, discard Status 真正执行cmd程序
	<-envCmd.Start()

	// Wait for goroutine to print everything cmd等待上面的程序执行完成
	<-doneChan
	// 检测备份文件是否存在
	_, err := os.Stat(fileName + ".dmp")
	if os.IsNotExist(err) {
		global.Log.Fatal("导出异常,程序退出")
		return
	}
	fmt.Println("\n导出结束，请查看当前路径下的文件 ->", fileName+".dmp")
}
