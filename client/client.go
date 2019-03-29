package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	url string
}

func (c *Client) Ping() {
	fmt.Println("making request to ", c.url)
	http.NewRequest("GET", c.url, nil)
}
