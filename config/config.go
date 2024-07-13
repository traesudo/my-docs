package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	data, err := ioutil.ReadFile("./config.yml")
	dir, err := os.Getwd()
	fmt.Printf("当前工作目录: %s\n", dir)
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	// 解析 YAML 文件
	err = yaml.Unmarshal(data, &Env)
	if err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}

	// 输出读取到的配置
	fmt.Printf("Database Config: %+v\n", Env.Database)
}

type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"password"`
}

type Config struct {
	Database DataBase `yaml:"database"`
}

var Env Config
