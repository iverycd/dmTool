package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
)

func ping(ip string) {
	cmd := exec.Command("ping", ip)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ping %s 失败, err:%s\n", ip, err)
		return
	}
	output, err := simplifiedchinese.GBK.NewDecoder().Bytes(out)
	if err != nil {
		fmt.Printf("编码 %+v 失败, err:%s\n", out, err)
		return
	}
	fmt.Printf("ping %s 成功, 返回信息:%s\n", ip, output)
}
func main() {
	ping("172.16.90.19")
}
