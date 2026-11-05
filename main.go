package main

import (
	"blog_server/core"
	"blog_server/flags"
	"blog_server/global"

	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogger()
	logrus.Warnf("配置文件读取成功，当前配置文件路径为：%s", flags.FlagOptions.File)
	logrus.Warningf("xxxxxx")
	logrus.Errorf("ddddddddd")
	logrus.Debugf("debug信息")
	logrus.Infof("info信息")
}
