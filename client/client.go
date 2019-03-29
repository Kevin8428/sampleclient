package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	URL     string
	context context.Context
}

func (c *Client) Ping(m map[string]interface{}) error {
	c.context = context.Background()
	context.WithValue(c.context, "foo", "bar")
	return c.DoRequest(m)
}

func (c *Client) DoRequest(m map[string]interface{}) error {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	mm := map[string]interface{}{}
	json.Unmarshal(buf.Bytes(), &m)
	if err != nil {
		return err
	}
	m = mm
	return err
}
