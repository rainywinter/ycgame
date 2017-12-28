package common

import (
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Conn represent a websocket connection
type Conn interface {
	Request() *http.Request // nil if it's a client
	ReadMessage() ([]byte, error)
	WriteMessage([]byte) error
	Close() error
	IP() string
	SetWriteDeadline(t time.Time)
	SetReadDeadline(t time.Time)
}

// NewConn create websocket connection by spec url
func NewConn(url string) (Conn, error) {
	dialer := websocket.Dialer{}
	c, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return &conn{ws: c}, nil
}

type conn struct {
	ws *websocket.Conn
	r  *http.Request
}

func (c *conn) SetWriteDeadline(t time.Time) {
	_ = c.ws.SetWriteDeadline(t)
}

func (c *conn) SetReadDeadline(t time.Time) {
	_ = c.ws.SetReadDeadline(t)
}

func (c *conn) IP() string {
	addr := c.ws.RemoteAddr()
	return addr.(*net.TCPAddr).IP.String()
}

func (c *conn) Request() *http.Request {
	return c.r
}

func (c *conn) ReadMessage() ([]byte, error) {
	_, message, err := c.ws.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (c *conn) WriteMessage(msg []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(NetworkTimeout))
	return c.ws.WriteMessage(websocket.BinaryMessage, msg)
}

func (c *conn) Close() error {
	return c.ws.Close()
}
