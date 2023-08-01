package config

import (
	"github.com/BurntSushi/toml"
)

// Config 对应配置文件结构
type Config struct {
	Port     int    `toml:"port"`
	ReportApi string `toml:"report_api"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		return c, err
	}
	return c, nil
}
