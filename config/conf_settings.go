package config

// Server  yaml文件中标签是server:的部分
type Server struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// Database yaml文件中标签是database:的部分
type Database struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"DbName"`
}

// Config yaml配置文件通过Unmarshal解析到这个结构体
type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}
