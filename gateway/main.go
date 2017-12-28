package main

import (
	"flag"
	"net/http"
	"rw/ycgame/common"
	"rw/ycgame/gateway/handler"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	http.HandleFunc("/ws", handler.HandleWS)
	glog.Info("==>gateway started: " + common.GatewayAddr)
	glog.Fatal(http.ListenAndServe(common.GatewayAddr, nil))
}
