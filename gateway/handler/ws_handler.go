package handler

import (
	"net/http"
	"rw/ycgame/common"
	"rw/ycgame/proto/pb"

	"github.com/golang/protobuf/proto"

	"github.com/golang/glog"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// HandleWS handle websocket conns
func HandleWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Error("upgrade err:", err)
		return
	}
	defer c.Close()

	for {
		mt, b, err := c.ReadMessage()
		if err != nil {
			glog.Error("read err:", err)
			break
		}
		keepalive := &pb.Keepalive{}
		err = proto.Unmarshal(b, keepalive)
		if err != nil {
			glog.Error("unmarshal err:", err)
			break
		}
		glog.Infof("recv: %s", common.Name(keepalive))

		err = c.WriteMessage(mt, b)
		if err != nil {
			glog.Error("write err:", err, "msg: ", common.Name(keepalive))
			break
		}
	}
}
