package websocket

import (
	"fmt"
	"job-sheduler/internal/infra/store"
	"job-sheduler/internal/model/types"

	"github.com/gorilla/websocket"
)

// new socker user client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// singelton instance for new connection
var poolInstance *Pool

// Connection Pool implementation for websockets
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan string
}


// function to follow singleton pattern to get pool
func GetPoolInstance() *Pool {
	if poolInstance == nil {
		poolInstance = &Pool{}
	}
	return poolInstance
}

// new pool provider
func (pool *Pool) NewPool() {

	pool.Register = make(chan *Client)
	pool.Unregister = make(chan *Client)
	pool.Clients = make(map[*Client]bool)
	pool.Broadcast = make(chan string)

}

// listner to manage activites on the websocket pool
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
		case <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {

				response := types.JobResponse{
					Jobs:   *store.GetStoreInstance().GetStore(),
					Length: len(*store.GetStoreInstance().GetStore()),
				}
				if err := client.Conn.WriteJSON(response); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

// function to keep the clinet to listen 
func (c *Client) Read() {
    defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
    }()

    response := types.JobResponse{
        Jobs:   *store.GetStoreInstance().GetStore(),
        Length: len(*store.GetStoreInstance().GetStore()),
    }

    if err := c.Conn.WriteJSON(response); err != nil {
        fmt.Println(err)
        return
    }

    for {
        _, _, err := c.Conn.ReadMessage()
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Message Received")
    }
}