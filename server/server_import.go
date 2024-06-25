//go:build imp

package server

import (
	"bufio"
	"dmTool/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-cmd/cmd"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os"
	"regexp"
	"strings"
)

// ExpImp 导入
func ExpImp() {
	var isConfirm string
	var bakSchemaName string
	// 是否覆盖，默认为空字符串
	isReplace := ""
	// 获取除了程序名称之外的所有参数,这里是获取备份文件的完整路径
	inputFile := os.Args[1:]
	// [用户交互部分]
	if len(inputFile) > 0 {
		// 检测文件是否存在
		_, err := os.Stat(inputFile[0])
		if os.IsNotExist(err) {
			global.Log.Fatal(fmt.Sprintf("文件%s 不存在，程序退出", inputFile[0]))
			return
		}
		fmt.Println("你输入的文件: ", inputFile[0])

		// 读取dmp文件头部3999字节，然后获取到版本号
		header, err := ReadFileHeader(inputFile[0], 3999) // 假设你想读取前5个字节
		if err != nil {
			log.Fatal(err)
		}
		version := FetchDmpVersion(header)
		if version != "" {
			fmt.Println("备份文件头版本号:", version)
		} else {
			fmt.Println("备份文件头版本号获取失败")
		}

		//fmt.Printf("目标主机-> %s \n目标数据库连接账号-> %s "+
		//	"\n请确认信息是否正确，输入\"YES\"继续,或者\"NO\"退出此程序\n", global.Config.Server.Host,
		//	global.Config.Server.User,
		//)
		colorStr := color.New()
		colorStr.Add(color.FgHiRed)
		colorStr.Printf("目标主机-> %s \n目标数据库连接账号-> %s "+
			"\n请确认信息是否正确，输入\"YES\"继续,或者\"NO\"退出此程序\n", global.Config.Server.Host,
			global.Config.Server.User)
		// 只有在终端输入YES之后才会导入
		fmt.Scanln(&isConfirm)
		if strings.ToUpper(isConfirm) != "YES" {
			return
		}
		// 输入是否覆盖导入
		for {
			colorStr.Printf("是否覆盖导入请输入YES或者NO\n")
			fmt.Scanln(&isConfirm)
			if strings.ToUpper(isConfirm) == "YES" {
				isReplace = "TABLE_EXISTS_ACTION=replace"
				break
			} else if strings.ToUpper(isConfirm) == "NO" {
				break
			} else {
				fmt.Println("无效输入，请只输入 'yes' 或 'no'")
			}
		}

	} else {
		global.Log.Fatal("[请输入备份文件的绝对路径; 程序退出!]")
	}
	// [导入前的分析部分]

	// 读取备份文件重定向到平面文件
	// Run foo and block waiting for it to exit
	dmCmd := fmt.Sprintf("dm_client\\dimp %s/\"%s\"@%s:%s file=%s show=y >analyze.log",
		global.Config.Server.User,
		global.Config.Server.Password,
		global.Config.Server.Host,
		global.Config.Server.Port,
		inputFile[0],
	)
	fmt.Println(dmCmd)
	c := cmd.NewCmd("cmd", "/C", dmCmd)
	s := <-c.Start()
	// 完成状态，true就是完成了
	if !s.Complete {
		os.Exit(0)
	}
	// 对analyze文件进行解析获取源的模式名
	file, err := os.Open("analyze.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建一个扫描器
	scanner := bufio.NewScanner(file)
	// 扫描每行
	for scanner.Scan() {
		line := scanner.Text()
		//将GBK编码的字符串转换为utf-8编码
		line, err := simplifiedchinese.GBK.NewDecoder().String(line)
		if err != nil {
			global.Log.Fatal(err)
		}
		// 正则表达式匹配 '模式TEST  含有 334 个表',获取到模式名
		re := regexp.MustCompile(`^模式(\S+)\s+含有`)
		if re == nil {
			global.Log.Fatal(err)
		}
		matches := re.FindStringSubmatch(line)
		// 切片长度大于1，说明这一行匹配到了
		if len(matches) > 1 {
			bakSchemaName = matches[1]
			fmt.Printf("schema = %s\n", bakSchemaName)
			break // 只查找第一行匹配即可，这里假设文件中只有一个匹配
		}
	}
	// 如果上面正则没有匹配到，抛出警告退出程序
	if bakSchemaName == "" {
		global.Log.Fatal("[Error messages: fetch schema failed!]")
	}
	// 处理错误
	if err := scanner.Err(); err != nil {
		global.Log.Fatal(err)
	}

	// [导入部分]

	// 拼接导入命令,isReplace默认为空字符串，否则为TABLE_EXISTS_ACTION=replace
	//findCmd := cmd.NewCmd("cmd", "/C", "dm_client\\dimp test/Gepoint@192.168.74.10 file=dexp.dmp table_exists_action=replace schemas=test") 导入示例
	dmCmd = fmt.Sprintf("dm_client\\dimp %s/\"%s\"@%s:%s file=%s log=%s.log LOG_WRITE=y dummy=y %s remap_schema=%s:%s",
		global.Config.Server.User,
		global.Config.Server.Password,
		global.Config.Server.Host,
		global.Config.Server.Port,
		inputFile[0],
		inputFile[0],
		isReplace,
		strings.ToUpper(bakSchemaName),
		strings.ToUpper(global.Config.Server.User),
	)
	fmt.Println(dmCmd)

	// 执行导入

	// Disable output buffering, enable streaming,设置cmd的输出方式
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options,创建带选项的cmd，并不是真正执行cmd
	envCmd := cmd.NewCmdOptions(cmdOptions, "cmd", "/C", dmCmd)

	// Print STDOUT and STDERR lines streaming from Cmd 实时输出cmd运行的命令行程序输出的内容
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

	// Run and wait for Cmd to return, discard Status 正在执行命令行程序
	<-envCmd.Start()

	// Wait for goroutine to print everything 等待上面的命令行程序在cmd运行完成
	<-doneChan

	global.Log.Info("导入已结束，请查看导入日志", inputFile[0]+".log")

}
