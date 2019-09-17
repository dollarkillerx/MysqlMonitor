/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:11 2019-09-17
 */
package logic

import (
	"MysqlMonitor/config"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/gemail"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os/exec"
)

func PingDb() {
	defer func() {
		if p := recover(); p != nil {
			clog.Println("数据库连接失败")
			emails := make([]string,0)
			emails = append(emails, config.Basis.App.Email)
			gemail.SendNifoLog(emails,"mysql宕机了","mysql宕机了")
			rebootMysql()  // 尝试重启mysql
			return
		}
	}()

	engine, e := xorm.NewEngine("mysql", config.Basis.Mysql.Dsn)
	if e != nil {
		panic(e)
	}

	e = engine.Ping()
	if e != nil {
		panic(e)
	}
}


func rebootMysql() {
	cmd := exec.Command("service", "mysql","restart")
	err := cmd.Start()
	if err != nil {
		clog.Println(err)
	}
	err = cmd.Wait()
}