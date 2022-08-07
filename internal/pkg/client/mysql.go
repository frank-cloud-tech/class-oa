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

// Package client
// @Description: 建立数据库连接
package client

import (
	"fmt"
	"github.com/frank-cloud-tech/class-oa/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var _iamDb *gorm.DB

// GetMysqlDB
// @Description: 不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// @return *gorm.DB:
func GetMysqlDB() *gorm.DB {
	return _iamDb
}

func init() {
	mysqlInfo := config.Config.GetMysql()
	var err error
	//connect to db and get db object for read or write option.
	_iamDb, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
				mysqlInfo.Username, mysqlInfo.Password, mysqlInfo.Address,
				mysqlInfo.Port, mysqlInfo.DBName, mysqlInfo.Timeout),
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",   // table name prefix, if "t_" table for `User` would be `t_users`
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
			DisableForeignKeyConstraintWhenMigrating: true,
		})

	if err != nil {
		panic("Failed to connect database, error=" + err.Error())
	}
	sqlDB, _ := _iamDb.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(mysqlInfo.MaxConn)     //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(mysqlInfo.MaxIdleConn) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数>20，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
