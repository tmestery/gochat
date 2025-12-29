package client

import "net/rpc"

type Message struct {
	Username string
	Body     string
}

type Client struct {
	rpc *rpc.Client
}

func New(addr string) (*Client, error) {
	c, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return nil, err
	}
	
	return &Client{rpc: c}, nil
}

func (c *Client) GetMessages(reply *[]Message) error {
	return c.rpc.Call("API.GetMessages", struct{}{}, reply)
}

func (c *Client) AddMessage(msg Message, reply *Message) error {
	return c.rpc.Call("API.AddMessage", msg, reply)
}