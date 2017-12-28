package data

import (
	"log"
	"runtime/debug"
	"rw/ycgame/common"
	"rw/ycgame/proto/pb"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/golang/glog"
)

// Player 玩家
type Player struct {
	ID int

	conn common.Conn

	recv chan proto.Message
	send chan proto.Message

	halt chan struct{}
}

func NewPlayer(id int) (*Player, error) {
	conn, err := common.NewConn("ws://" + common.GatewayAddr + "/ws")
	if err != nil {
		return nil, err
	}
	return &Player{
		ID:   id,
		conn: conn,
		recv: make(chan proto.Message, 100),
		send: make(chan proto.Message, 100),
		halt: make(chan struct{}),
	}, nil
}

func (p *Player) Loop() {
	glog.Infof("player: %d enter loop", p.ID)
	defer func() {
		glog.Infof("player: %d exit loop")
	}()
	defer func() {
		if r := recover(); r != nil {
			log.Println(r, debug.Stack())
		}
	}()

	defer func() {
		p.conn.Close()
	}()

	errCh := make(chan error, 100)

	// begin read
	go func() {
		for {
			b, err := p.conn.ReadMessage()
			if err != nil {
				errCh <- err
				return
			}

			keepAlive := &pb.Keepalive{}
			proto.Unmarshal(b, keepAlive)
			p.recv <- keepAlive
		}
	}()

	// write
	go func() {
		for msg := range p.send {
			b, err := proto.Marshal(msg)
			if err != nil {
				errCh <- err
				return
			}
			err = p.conn.WriteMessage(b)
			if err != nil {
				errCh <- err
				return
			}
			glog.Infof("player: %d send %v", p.ID, common.Name(msg))
		}
	}()

	keepaliveTicker := time.NewTicker(time.Second * 1)

	for {
		select {
		case msg := <-p.recv:
			glog.Infof("player: %d recv: %v", p.ID, common.Name(msg))

		case <-keepaliveTicker.C:
			p.KeepAlive()

		case err := <-errCh:
			log.Printf("catch err:%v,exit...", err)
			return

		case <-p.halt:
			return
		}
	}
}

func (p *Player) Send(msg proto.Message) {
	select {
	case p.send <- msg:

	default: // blocked
		log.Println("send queue block, shutdown")
		close(p.send)
	}
}

func (p *Player) KeepAlive() {
	p.Send(&pb.Keepalive{})
}

func (p *Player) Halt() {
	close(p.halt)
}
