package common

import (
	"time"
)

var (

	// GatewayAddr 游戏网关地址
	GatewayAddr = "127.0.0.1:8080"

	// GameAddr 游戏服务器地址
	GameAddr = "127.0.0.1:8081"

	// NetworkTimeout 网络io超时时间
	NetworkTimeout = time.Second * 5
)
