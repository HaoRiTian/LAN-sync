package ws

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/gorilla/websocket"
    "net/http"
    "sync"
)

var (
    once      = sync.Once{}
    wsManager *WebsocketManager
)

type WebsocketManager struct {
    Clients              map[string]*Client
    clientCount          uint
    Lock                 sync.Mutex
    Register, UnRegister chan *Client
    BroadCastMessage     chan []byte
}

func NewWsManager() *WebsocketManager {
    once.Do(func() {
        wsManager = &WebsocketManager{
            Clients:          make(map[string]*Client),
            Register:         make(chan *Client, 128),
            UnRegister:       make(chan *Client, 128),
            BroadCastMessage: make(chan []byte, 1024),
            clientCount:      0,
        }
    })
    return wsManager
}

func (m *WebsocketManager) Run() {
    go m.Start()
    go m.SendAllService()
}

func (m *WebsocketManager) Info() map[string]interface{} {
    managerInfo := make(map[string]interface{})
    managerInfo["clientLen"] = m.LenClient()
    managerInfo["chanRegisterLen"] = len(m.Register)
    managerInfo["chanUnregisterLen"] = len(m.UnRegister)
    managerInfo["chanBroadCastMessageLen"] = len(m.BroadCastMessage)
    return managerInfo
}

type BroadCastMessageData struct {
    Message []byte
}

func (m *WebsocketManager) Start() {
    for {
        select {
        case client := <-m.Register:
            m.Lock.Lock()

            m.Clients[client.Id] = client
            m.clientCount += 1
            m.Lock.Unlock()

        case client := <-m.UnRegister:
            m.Lock.Lock()
            if _, ok := m.Clients[client.Id]; ok {
                close(client.ToBeSentMessage)
                close(client.ToBeReadMessage)
                delete(m.Clients, client.Id)
                m.clientCount -= 1
            }
            m.Lock.Unlock()
        }
    }
}

func (m *WebsocketManager) WsClientHandler(ctx *gin.Context) {
    upGrader := websocket.Upgrader{
        // cross origin domain
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
        // 处理 Sec-WebSocket-Protocol Header
        Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
    }

    conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
    if err != nil {
        return
    }

    client := &Client{
        Id:              uuid.New().String(),
        Socket:          conn,
        ToBeSentMessage: make(chan []byte, 1024),
        ToBeReadMessage: make(chan []byte, 1024),
    }
    m.Register <- client
    go client.Read()
    go client.Write()

    go client.heartbeat()
}

func (m *WebsocketManager) SendAllService() {
    for {
        select {
        case data := <-m.BroadCastMessage:
            for _, client := range m.Clients {
                client.ToBeSentMessage <- data
            }
        }
    }
}

func (m *WebsocketManager) SendAll(message []byte) {
    m.BroadCastMessage <- message
}

func (m *WebsocketManager) LenClient() uint {
    return m.clientCount
}
