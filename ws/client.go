package ws

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "time"
)

type Client struct {
    Id     string
    Socket *websocket.Conn
    // 待发送消息
    ToBeSentMessage chan []byte
    // 待读取消息 接收的消息放入此通道
    ToBeReadMessage chan []byte
}

// 读信息，从 websocket 连接直接读取数据
func (c *Client) Read() {
    defer func() {
        wsManager.UnRegister <- c
        if err := c.Socket.Close(); err != nil {
        }
    }()

    for {
        messageType, message, err := c.Socket.ReadMessage()
        if err != nil || messageType == websocket.CloseMessage {
            break
        }
        c.ToBeReadMessage <- message
    }
}

// 写信息，从管道中读取数据写入 websocket 连接
func (c *Client) Write() {
    defer func() {
        wsManager.UnRegister <- c
        if err := c.Socket.Close(); err != nil {
        }
    }()

    for {
        select {
        case message, ok := <-c.ToBeSentMessage:
            if !ok {
                _ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            err := c.Socket.WriteMessage(websocket.BinaryMessage, message)
            if err != nil {
                return
            }
        }
    }
}

func (c *Client) heartbeat() {
    defer func() {
        if recover() != nil {
        }
    }()

    for {
        msg, err := json.Marshal(time.Now().Format("2006-01-02 15:04:05"))
        if err != nil {
        }
        c.ToBeSentMessage <- msg
        time.Sleep(time.Second * 60)
    }
}
