package main

import (
	"fmt"
	"minigrep/proc"
	"os"
)

func main() {

	//命令行提示
	promptTips := "[minigrep] Usage: minigrep [QUERY] [FILEPATH]"

	//从用户输入中获取参数
	config := proc.Config{}
	err := config.Build(os.Args)
	if err != nil {
		fmt.Println(promptTips)
		return
	}

	//根据参数执行任务
	err = proc.Run(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
