package main

import (
	"flag"
	"net"

	"rw/ycgame/common"
	"rw/ycgame/proto/pb"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, p *pb.Ping) (*pb.Pong, error) {
	glog.Info("recv: ", p.Msg)
	return &pb.Pong{Msg: "i am game"}, nil
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", common.GameAddr)
	if err != nil {
		glog.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterPingPongServer(s, &server{})

	glog.Infof("start game server: %v", common.GameAddr)
	s.Serve(l)
}
