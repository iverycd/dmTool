//go:build imp

package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"vastTool/global"
)

// ExpImp 导入
func (connect *Connect) ExpImp() {
	var isConfirm string
	// 获取除了程序名称之外的所有参数,这里是获取备份文件的完整路径
	inputFile := os.Args[1:]
	if len(inputFile) > 0 {
		fmt.Println("Your input file name is: ", inputFile[0])
		fmt.Printf("Target host-> %s Database name-> %s "+
			"\nPlease confirm whether the information is correct and input \"YES\" for continue!\n", global.Config.Server.Host,
			global.Config.Database.DbName,
		)
		// 只有在终端输入YES之后才会导入
		fmt.Scanln(&isConfirm)
		if strings.ToUpper(isConfirm) == "YES" {
			fmt.Println("You input ", isConfirm)
		} else {
			return
		}
	} else {
		global.Log.Fatal("[Error messages: please input backup file absolute path; Program Exit!]")
	}
	// 建立ssh会话
	sftpClient, session := connect.InitSsh()
	// 上传备份
	// 打开本地的dump.sql文件
	localFile, err := os.Open(inputFile[0])
	if err != nil {
		global.Log.Fatal(err)
	}
	defer localFile.Close()
	// 创建服务器上的目标文件
	remotePath := "/tmp/" // 服务端上用来放数据库备份的目录
	_, fileName := filepath.Split(inputFile[0])
	remoteFile, err := sftpClient.Create(remotePath + fileName)
	if err != nil {
		global.Log.Fatal(err)
	}
	defer remoteFile.Close()
	// 本地文件的大小信息
	localFileInfo, _ := localFile.Stat()
	localFileSize := localFileInfo.Size()
	// 实时显示上传进度
	go func() {
		for {
			remoteFileInfo, _ := remoteFile.Stat()
			remoteFileSize := remoteFileInfo.Size()
			progress := float64(remoteFileSize) / float64(localFileSize) * 100
			fmt.Printf("\rFile %s  is Uploading ... Current Progress->  %.2f%%", fileName, progress)
			if remoteFileSize == localFileSize {
				break
			}
		}
	}()
	// 将本地文件内容复制到服务器上的目标文件
	//bytes, err := ioutil.ReadAll(localFile) // 读的os打开的文件句柄
	//bytes, err := os.ReadFile(inputFile[0]) // 读的文件绝对路径
	//if err != nil {
	//	global.Log.Fatal(err)
	//}
	//_, err = remoteFile.Write(bytes)
	//if err != nil {
	//	global.Log.Fatal(err)
	//}
	// ReadFrom是全速上传
	_, err = remoteFile.ReadFrom(localFile)
	if err != nil {
		global.Log.Fatal(err)
	}
	// 检查文件是否成功上传
	info, err := sftpClient.Stat(remotePath + fileName)
	if err != nil {
		global.Log.Fatal(err)
	}
	if info.Size() == 0 {
		global.Log.Fatal(" \nFile upload failed")
	} else {
		fmt.Println(" \nFile uploaded successfully")
	}
	// 执行导入命令
	global.Log.Println("begin import to database <- " + remotePath + fileName)
	cmd := fmt.Sprintf("su - vastbase -c \"vsql -U%s -W%s -p%s %s < %s\"",
		global.Config.Database.User,
		global.Config.Database.Password,
		global.Config.Database.Port,
		global.Config.Database.DbName,
		remotePath+fileName,
	)
	global.Log.Println(cmd)
	// 运行导入
	err = session.Run(cmd)
	defer session.Close()
	if err != nil {
		global.Log.Fatalf("Failed to execute command '%s': %s", cmd, err)
	}
	global.Log.Info("Import finish to vastbase")
}
