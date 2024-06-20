package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	//str := "模式TEST  含有 334 个表，分别为："
	//versions := "i am Binary,   Version: 5.7.27-log (i am Version description ). started"
	versions := "xx模式TEST  含有 334 个表，分别为："
	//versionre := regexp.MustCompile(`模式(.+?)  含有`)
	versionre := regexp.MustCompile(`^模式(\S+)\s+含有`)
	if versionre == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	matches := versionre.FindStringSubmatch(versions)
	fmt.Println("len matches:", len(matches))
	fmt.Println("result1 = ", matches)

	//read, err := ReverseRead("test.dmp.log", 10)
	//if err != nil {
	//	return
	//}
	//for _, content := range read {
	//	//将GBK编码的字符串转换为utf-8编码
	//	output, err := simplifiedchinese.GBK.NewDecoder().String(content)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(output)
	//}

	logPath := "test.dmp.log"
	lines, err := readLastLines(logPath, 10)
	if err != nil {
		fmt.Println("Error reading last lines:", err)
		return
	}

	for _, line := range lines {
		fmt.Print(line)
	}

}

func readLastLines(path string, n int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 创建缓冲读取器
	reader := bufio.NewReader(file)

	// 读取文件从末尾开始的n行
	var lines []string
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // 到达文件开始或出错
		}
		lines = append([]string{line}, lines...) // 添加到切片头部
	}

	return lines, nil
}

func ReverseRead(name string, lineNum uint) ([]string, error) {
	//打开文件
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//获取文件大小
	fs, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fs.Size()

	var offset int64 = -1   //偏移量，初始化为-1，若为0则会读到EOF
	char := make([]byte, 1) //用于读取单个字节
	lineStr := ""           //存放一行的数据
	buff := make([]string, 0, 100)
	for (-offset) <= fileSize {
		//通过Seek函数从末尾移动游标然后每次读取一个字节
		file.Seek(offset, io.SeekEnd)
		_, err := file.Read(char)
		if err != nil {
			return buff, err
		}
		if char[0] == '\n' {
			offset--  //windows跳过'\r'
			lineNum-- //到此读取完一行
			buff = append(buff, lineStr)
			lineStr = ""
			if lineNum == 0 {
				return buff, nil
			}
		} else {
			lineStr = string(char) + lineStr
		}
		offset--
	}
	buff = append(buff, lineStr)
	return buff, nil
}
