package options

import (
	"net/url"
)

type ClientOptions struct {
	*url.URL
}

func Client() *ClientOptions {
	return &ClientOptions{}
}

func (o *ClientOptions) ApplyURI(uri string) *ClientOptions {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil
	}
	o.URL = u
	return o
}
