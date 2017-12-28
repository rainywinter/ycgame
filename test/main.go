package main

import (
	"flag"
	"rw/ycgame/common"
	"rw/ycgame/proto/pb"

	"github.com/golang/glog"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(common.GameAddr, grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPingPongClient(conn)

	p, err := c.Greet(context.Background(), &pb.Ping{Msg: "i am from test"})
	if err != nil {
		glog.Fatalf("ping err: %v", err)
	}
	glog.Infof("pong: %v", p.Msg)
}
