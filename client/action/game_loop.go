package action

import (
	"rw/ycgame/client/data"

	"github.com/golang/glog"
)

var shutdown chan struct{}

func init() {
	shutdown = make(chan struct{})
}
func Run() (err error) {

	p, err := data.NewPlayer(1)
	if err != nil {
		return err
	}
	go p.Loop()

	<-shutdown
	glog.Info("exit loop")
	return nil
}

// Shutdown shutdown for loop
func Shutdown() {
	close(shutdown)
}
