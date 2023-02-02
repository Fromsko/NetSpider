package main

import (
	app "spider/code"
	"spider/util"
)

func main() {
	/*
		Start application
			+ 接受输入
				+ 校验输入信息
			+ 爬虫 ==> (默认 './res')
			+ 日志 ==> (默认 './res/log/*.log')
		End
	*/
	downUrl := util.GetInput()
	app.Fetch(downUrl, "./res")
}
