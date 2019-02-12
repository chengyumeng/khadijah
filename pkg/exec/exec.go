package exec

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/chengyumeng/khadijah/pkg/config"
	utillog "github.com/chengyumeng/khadijah/pkg/utils/log"

	"github.com/gorilla/websocket"
	"github.com/pkg/term"
)

var (
	logger = utillog.NewAppLogger("pkg/exec")
)

// shell command type
const (
	STDIN   = "stdin"
	CONNECT = "connect"
)

// SocketShell 并不是 ssh 的全称，ssh的全称是 Sentry Socket
// 这里这样定义的目的只是描述 shell 的实现形式
type SocketShell struct {
	Connection *websocket.Conn
	Command    string
	Exit       bool
}

// Message is the socket message of ssh
type Message struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

// NewSocketShell is the function to init a new socket shell
func NewSocketShell() *SocketShell {
	ssh := new(SocketShell)
	u := fmt.Sprintf("%s/api/v1/clientool/exec", config.GlobalOption.System.WebsocketURL)
	var h = make(http.Header)
	h.Set("Authorization", "Bearer "+config.GlobalOption.Token)
	c, _, err := websocket.DefaultDialer.Dial(u, h)
	if err != nil {
		logger.Errorln(err)
		c.Close()
	}
	ssh.Connection = c
	return ssh
}

// Connect is the function to connect to remote k8s pod
func (s *SocketShell) Connect(option Option) error {
	sendData := Message{CONNECT, option}
	data, err := json.Marshal(sendData)
	if err != nil {
		return err
	}
	return s.Connection.WriteMessage(websocket.TextMessage, data)
}

// Listen is the function to listen to remote k8s pod
func (s *SocketShell) Listen() {
	writer := bufio.NewWriter(os.Stdout)
	for s.Exit == false {
		_, message, err := s.Connection.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}
		writer.WriteString(string(message))
		fmt.Print(string(message))
		if string(message) == "\r\nexit\r\n" {
			s.Exit = true
			logger.Infoln("Safely close the connection with the Enter key")
		}
	}
}

// StdinSend is the function to send message from stdin to remote k8s pod
func (s *SocketShell) StdinSend() {
	t, _ := term.Open("/dev/tty")
	defer func() {
		t.Restore()
		t.Close()
	}()
	term.RawMode(t)

	bytes := make([]byte, 1)

	for s.Exit != true {
		t.Read(bytes)
		s.Command = string(bytes)
		sendData := Message{STDIN, s.Command}
		data, err := json.Marshal(sendData)
		if err != nil {
			fmt.Println(err)
		}
		s.Connection.WriteMessage(websocket.TextMessage, data)
	}
}
