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

type conf struct {
	Config config `yaml:"config"`
}

type config struct {
	LogDir string `yaml:"Log_dir"`
	OA     OA     `yaml:"oa"`
	Mysql  Mysql  `yaml:"mysql"`
}

type OA struct {
	Port int `yaml:"port"`
}

type Mysql struct {
	Address     string `yaml:"address"`
	Port        int    `yaml:"port"`
	DBName      string `yaml:"dbname"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxConn     int    `yaml:"max_conn"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
	Timeout     string `yaml:"timeout"`
}

func (c *conf) GetMysql() *Mysql {
	return &c.Config.Mysql
}

func (c *conf) GetLogDir() string {
	return c.Config.LogDir
}

func (c *conf) GetOAPort() int {
	return c.Config.OA.Port
}
