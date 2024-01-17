package config

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	DateTimeFormat = "2006-01-02 15:04:05.999 -07:00"
	DateFomart     = "2006-01-02"
	TimeFormat     = "15:04:05"
	UpGrade        = websocket.Upgrader{ // 防止跨域站点伪造请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
