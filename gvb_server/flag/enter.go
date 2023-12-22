package flag

import sysflag "flag"

type Option struct {
	Version bool
	DB      bool
}

// Parse 解析命令行参数
func Parse() Option {
	version := sysflag.Bool("v", false, "Version")
	db := sysflag.Bool("db", false, "Initialize database")
	// 解析命令行参数, 写入注册的flag中
	sysflag.Parse()
	return Option{
		Version: *version,
		DB:      *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return false
	}
	if option.Version {
		return false
	}
	return true
}

// SwitchOption 根据命令行执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
	if option.Version {
		PVersion()
	}
}
