package main

import (
	"flag"
	"rw/ycgame/client/action"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Fatal(action.Run())
}
