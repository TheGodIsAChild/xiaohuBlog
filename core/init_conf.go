package core

import (
	"blog_server/conf"
	"blog_server/flags"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConf() *conf.Config {
	path := flags.FlagOptions.File
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("yaml文件读取错误【文件不存在】：%s", err))
	}
	config := &conf.Config{}
	// 映射yaml文件到结构体
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml文件读取错误【文件格式错误】：%s", err))
	}
	return config
}
