/*
Copyright © 2022 frank-cloud-tech

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package config
// @Description:
package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

import (
	"flag"
	"testing"
)

var env = flag.String("conf", "dev", "配置文件")

// Config 项目配置
var Config conf

//
//  init
//  @Description: 加载并读配置文件
//
func init() {
	Config = loadConfig()
	testing.Init()
	flag.Parse()
}

//
//  loadConfig
//  @Description: 加载项目的配置文件, 默认的配置文件和项目的可执行文件同级
//  @return *Config
//
func loadConfig() conf {
	configFile, err := ioutil.ReadFile("./config-" + *env + ".yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
	return Config
}
