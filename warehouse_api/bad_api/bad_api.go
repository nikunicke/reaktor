package bad_api

import (
	"net/http"
	"time"

	"github.com/nikunicke/reaktor/warehouse_api/bad_api/options"
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
	url := c.o.URL.String() + t
	// fmt.Println(url)
	resp, err = c.c.Get(url)
	return resp, err
}
