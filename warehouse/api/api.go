package api

import (
	"net/http"
	"time"

	"github.com/nikunicke/reaktor/warehouse"
	"github.com/nikunicke/reaktor/warehouse/api/options"
)

type Client struct {
	c *http.Client
	o *options.ClientOptions
}

func NewClient(opts ...*options.ClientOptions) (*Client, error) {
	cli := &Client{
		c: &http.Client{},
		o: &options.ClientOptions{},
	}
	cli.c.Timeout = time.Second * 5
	for _, o := range opts {
		if o == nil {
			continue
		}
		if o.URL != nil {
			cli.o.URL = o.URL
		}
	}
	return cli, nil
}
func (c *Client) Get(t string) (resp *http.Response, err error) {
	// fmt.Println("Requesting:", t)
	url := c.o.URL.String() + t
	resp, err = c.c.Get(url)
	e := resp.Header["X-Error-Modes-Active"]
	if e[0] != "" {
		return nil, warehouse.Error("Failed to request data from endpoint: " + t)
	}
	return resp, nil
}
