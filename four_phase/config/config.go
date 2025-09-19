package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database  Database `yaml:"db"`
	SecretKey string   `yaml:"key"`
}

type Database struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func LoadConfig(filename string) (*Config, error) {
	// 读取文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %v", err)
	}

	// 解析YAML到结构体
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}
