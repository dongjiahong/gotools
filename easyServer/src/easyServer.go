package es

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// EasyServer ..
type EasyServer struct {
	Config EasyConf `json:"easy_conf"`
	EasyLogger
}

// EasyConf ..
type EasyConf struct {
	//ConfFile string `json:"conf_file"`
	LogPath string `json:"log_path"`
}

// EasyLogger ..
type EasyLogger struct {
	l *log.Logger
}

// ReadConfig 读取json格式的配置文件，并返回EasyConf结构体指针
func ReadConfig(config string) *EasyConf {
	// 加载默认配置文件
	conf, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Println("read config err: ", err)
		return nil
	}
	var easyConf EasyConf
	err = json.Unmarshal(conf, easyConf)
	if err != nil {
		fmt.Println("Unmarshal conf err: ", err)
		return nil
	}
	return &easyConf
}

// InitLogger 初始化logger 参数为log的名字和log的前缀
func (s *EasyServer) InitLogger(logFileName string, logPreffix string) error {
	logFile := s.Config.LogPath + logFileName
	file, err := os.Create(logFile)
	if err != nil {
		return errors.New("create logfile err: " + err.Error())
	}
	defer file.Close()
	s.l = log.New(file, logPreffix, log.Llongfile)
	return nil
}

// ReadConfigFromFile 从文件中读取json格式的文件配置
func (s *EasyServer) ReadConfigFromFile(configFile string) error {
	if len(configFile) == 0 {
		panic("can't find config file")
	}
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.New("read config err: " + err.Error())
	}
	err = json.Unmarshal(conf, &s.Config)
	if err != nil {
		return errors.New("unmarshal config err: " + err.Error())
	}
	return nil
}

// NewServer 创建一个新的server
func (s *EasyServer) NewServer(configFile string, logFileName string, logPreffix string) error {
	err := s.ReadConfigFromFile(configFile)
	if err != nil {
		return errors.New("ReadConfigFromFile err: " + err.Error())
	}
	err = s.InitLogger(logFileName, logPreffix)
	if err != nil {
		return errors.New("InitLogger err: " + err.Error())
	}
	return nil
}
