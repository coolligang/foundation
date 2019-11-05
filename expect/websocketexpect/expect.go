package websocketexpect

import (
	"errors"
	"github.com/gorilla/websocket"
	"gopkg.in/gavv/httpexpect.v2"
	"net/http"
	"testing"
	"time"
)

type webSocketExpect struct {
	Conn *websocket.Conn
}

func NewWebSocketExpect(baseURL string) *webSocketExpect {
	d := websocket.Dialer{}
	conn, resp, err := d.Dial(baseURL, nil)
	if err != nil {
		return nil
	}
	if got, want := resp.StatusCode, http.StatusSwitchingProtocols; got != want {
		return nil
	}
	item := new(webSocketExpect)
	item.Conn = conn
	return item
}

func (expect webSocketExpect) ReadJSONExpect(t *testing.T, timeout time.Duration) *httpexpect.Value {
	if expect.Conn == nil {
		t.Fatal("expect.Conn is nil")
	}
	var message interface{}
	_ = expect.Conn.SetReadDeadline(time.Now().Add(timeout))
	err := expect.Conn.ReadJSON(&message)
	if err != nil {
		t.Fatal(err)
	}
	reporter := httpexpect.NewRequireReporter(t)
	return httpexpect.NewValue(reporter, message)
}

func (expect webSocketExpect) Close() error {
	if expect.Conn == nil {
		return errors.New("expect.Conn is nil")
	}
	return expect.Conn.Close()
}
