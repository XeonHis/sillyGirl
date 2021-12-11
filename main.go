package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/cdle/sillyGirl/core"
)

func main() {
	go core.RunServer()
	logs.Info("傻妞用不了了？关注频道 https://t.me/kczz2021 获取最新消息。")
	fmt.Println("你可以在这里输入命令：")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		core.Senders <- &core.Faker{
			Type:    "terminal",
			Message: string(data),
		}
	}
}
