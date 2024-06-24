package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

// ReadFileHeader 读取文件的前N个字节
func ReadFileHeader(filename string, n int64) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件的前N个字节
	header := make([]byte, n)
	_, err = file.Read(header)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func FetchDmpVersion(header []byte) string {
	// 用来存取dmp文件去除数字0的部分，这部分字符是不可见字符
	var lines []uint8
	for _, content := range header {
		if content != 0 {
			lines = append(lines, content)
		}
	}

	// 第一次 正则表达式匹配 03134284172-20240321-222308-200935 获取dmp文件导出的版本
	re := regexp.MustCompile(`(.*)--([0-9]+)-([0-9]+)-([0-9]+)-([0-9]+)`)
	matches := re.FindStringSubmatch(string(lines))
	// 切片长度大于1，说明这一行匹配到了
	if len(matches) > 0 {
		for _, match := range matches {
			// 只匹配字符串长度为8的，比如20240321就是版本号
			if len(match) == 8 {
				return match
			}
		}
	}
	// 如果上面第一次没匹配到，紧接着 第二次 正则表达式匹配 1-2-2-21.05.13-139380-10006-ENT  Pack5
	re = regexp.MustCompile(`([0-9]+)-([0-9]+)-([0-9]+)-([0-9{2}.]+)(.*)`)
	matches = re.FindStringSubmatch(string(lines))
	if len(matches) > 0 {
		for _, match := range matches {
			// 只匹配字符串长度为8的，比如21.05.13就是版本号
			if len(match) == 8 {
				return match
			}
		}
	}
	// 啥都没有匹配到，返回空字符串
	return ""
}

func main() {
	// 读取dmp文件头部3999字节
	header, err := ReadFileHeader("test.dmp", 3999) // 假设你想读取前5个字节
	if err != nil {
		log.Fatal(err)
	}
	version := FetchDmpVersion(header)
	if version != "" {
		fmt.Println(version)
	}
}
