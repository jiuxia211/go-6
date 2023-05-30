package config

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode      string
	HttpPort     string
	DbUser       string
	DbPassWord   string
	DbHost       string
	DbName       string
	DbPort       string
	Path         string
	Domain1      string
	Domain2      string
	Domain3      string
	Version      string
	JwtSecret    string
	GrpcAddress1 string
	GrpcAddress2 string
	EtcdAddress  string
	ReAddress    string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	Path = strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort,
		")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	Domain1 = file.Section("server").Key("Domain1").String()
	Domain2 = file.Section("server").Key("Domain2").String()
	Domain3 = file.Section("server").Key("Domain3").String()
	Version = file.Section("server").Key("Version").String()
	JwtSecret = file.Section("server").Key("JwtSecret").String()
	GrpcAddress1 = file.Section("server").Key("GrpcAddress1").String()
	GrpcAddress2 = file.Section("server").Key("GrpcAddress2").String()
	EtcdAddress = file.Section("etcd").Key("EtcdAddress").String()
	ReAddress = file.Section("redis").Key("ReAddress").String()
}
