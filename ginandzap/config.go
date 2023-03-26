package main

// Config 整个项目的配置
type Config struct {
	Mode       string `json:"mode"`
	Port       int    `json:"port"`
	*LogConfig `json:"log"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// Conf 全局配置变量
var Conf = new(Config)

func Init() error {
	Conf = &Config{
		Mode: "dev",
		Port: 9000,
		LogConfig: &LogConfig{
			Level:      "debug",
			Filename:   "a.txt",
			MaxSize:    200,
			MaxAge:     7,
			MaxBackups: 10,
		},
	}
	return nil
}
