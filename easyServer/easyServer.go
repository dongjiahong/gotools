package easyServer

import "io/ioutil"
import "fmt"
import "encoding/json"

// EasyServer ..
type EasyServer struct {
	Config *EasyConf   `json:"easy_conf"`
	Logger *EasyLogger `json:"easy_logger"`
}

// EasyConf ..
type EasyConf struct {
	ConfFile string `json:"conf_file"`
	LogPath  string `json:"log_path"`
}

// EasyLogger ..
type EasyLogger struct {
}

func (s *EasyServer) initLogger() {}

func (s *EasyServer) readConfFromFile(configFile string) *EasyConf {
	if len(configFile) == 0 {
		panic("can't find config file")
	}
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("read config err: ", err)
	}
	var easyConf EasyConf
	err = json.Unmarshal([]byte(conf), easyConf)
	return &easyConf
}

// NewServer 创建一个server服务
func (s *EasyServer) NewServer(logger *EasyLogger, config *EasyConf) *EasyServer {
}
