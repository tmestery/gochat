package client

import (
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
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

func (c *Client) GetDB(reply *[]Item) error {
	return c.rpc.Call("API.getMessageList", "", reply)
}

func (c *Client) GetByName(item Item, reply *Item) error {
	return c.rpc.Call("API.GetByName", item, reply)
}

func (c *Client) AddItem(item Item, reply *Item) error {
	return c.rpc.Call("API.AddItem", item, reply)
}

func (c *Client) EditItem(item Item, reply *Item) error {
	return c.rpc.Call("API.EditItem", item, reply)
}

func (c *Client) DeleteItem(item Item, reply *Item) error {
	return c.rpc.Call("API.DeleteItem", item, reply)
}