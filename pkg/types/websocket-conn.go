package types

import (
	coder "github.com/coder/websocket"
	"github.com/gorilla/websocket"
)

type WebSocketConn struct {
	EventEmitter

	*websocket.Conn
	Coder         *coder.Conn
	RemoteAddress string
}

func (t *WebSocketConn) Close() error {
	defer t.Emit("close")
	if t.Coder != nil {
		return t.Coder.CloseNow()
	}
	return t.Conn.Close()
}
