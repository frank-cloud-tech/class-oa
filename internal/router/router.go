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

// Package router
// @Description:
package router

import (
	"github.com/gin-gonic/gin"
)

var OARouter = gin.Default()
var v1 = OARouter.Group("/api/v1")

func RegisterOA(method, path string, router func(c *gin.Context)) {
	switch method {
	case "GET":
		v1.GET(path, router)
	case "POST":
		v1.POST(path, router)
	case "PUT":
		v1.PUT(path, router)
	case "DELETE":
		v1.DELETE(path, router)
	case "HEAD":
		v1.DELETE(path, router)
	case "CONNECT":
		v1.DELETE(path, router)
	case "OPTIONS":
		v1.DELETE(path, router)
	case "TRACE":
		v1.DELETE(path, router)
	}
}
