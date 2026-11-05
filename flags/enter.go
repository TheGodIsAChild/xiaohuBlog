package flags

// 解析命令行参数工具库
import "flag"

type Options struct {
	File    string
	DB      bool
	Version bool
	LogPath string
	AppName string
}

var FlagOptions = new(Options)

func Parse() {
	flag.StringVar(&FlagOptions.File, "f", "setting.yaml", "配置文件路径")
	flag.StringVar(&FlagOptions.LogPath, "log-path", "./logs", "日志文件路径")
	flag.StringVar(&FlagOptions.AppName, "app-name", "blog", "应用名称")
	flag.BoolVar(&FlagOptions.DB, "db", false, "是否初始化数据库")
	flag.BoolVar(&FlagOptions.Version, "v", false, "显示版本号")
	flag.Parse()
}
