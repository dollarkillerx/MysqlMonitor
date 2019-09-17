/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:04 2019-09-17
 */
package main

import (
	"MysqlMonitor/config"
	"MysqlMonitor/logic"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/robfig/cron"
	"io/ioutil"
)

func init() {
	b, e := easyutils.PathExists("config.yml")
	if e != nil || b == false {
		e := ioutil.WriteFile("config.yml", []byte(conf), 00666)
		if e != nil {
			panic("配置文件创建失败")
		}
		panic("请填写配置文件!")
	}
}

func main() {
	go taskMan()
	clog.Println("服务器正常启动")
	select {}
}


// 定时任务
func taskMan() {
	c := cron.New()
	e := c.AddFunc(config.Basis.App.Corn, logic.PingDb)
	if e != nil {
		panic(e.Error())
	}
	c.Start()
}


var conf = `
# 通用配置文件

app:
  corn: "0 */10 * * * *"
  email: "dollar@dolalr.com"

mysql:
  dsn: "root:mlbj@(127.0.0.1:3306)/show?charset=utf8"
`