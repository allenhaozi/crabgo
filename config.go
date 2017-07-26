package crabgo

import (
	//"github.com/Allenhaozi/crabgo/config"
	//"io/ioutil"
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	CConfig *CrabConfig
)

type CrabConfig struct {
	RunMode    string
	AppName    string
	AppPath    string
	CfgContent map[string]string
}

func init() {
	Dump("allen")
	CConfig = NewConfig()
	CConfig.ParseConfigContent()
}

func NewConfig() *CrabConfig {

	return &CrabConfig{
		CfgContent: make(map[string]string),
	}
}

func (self *CrabConfig) ParseConfigContent() {

	file, err := os.Open("conf/app.conf")
	if err != nil {
		Dump(err)
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			Dump(err)
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			Dump(line)
		}
	}
	if err != nil {
		Dump(err)
	} else {
		Dump(file)
	}
}

func (self *CrabConfig) GetConfig(key string) string {
	if data, ok := self.CfgContent[key]; ok {
		return data
	} else {
		return ""
	}
}
