/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:05 2019-09-17
 */
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type basisConf struct {
	App struct {
		Corn       string `yaml:"corn"`
		Email string `yaml:"email"`

	}
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}

}

var (
	Basis *basisConf
)

func init() {
	Basis = &basisConf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, Basis)
	if e != nil {
		panic(e.Error())
	}
}

