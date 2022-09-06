// Code generated by hertz generator.

package main

import (
	handler "hzminioupload/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...

	r.POST("/upload", handler.Upload)
	r.POST("/getuploadurl", handler.GetUploadUrl)

}
