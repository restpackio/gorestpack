// This package provides easy access to http://restpack.io API services from Go applications
package gorestpack

import (
	"github.com/eknkc/request"
)

type client struct {
	httpClient  request.Client
	accessToken string
	basePath    string
}

func (me *client) do(method string, path string) request.Session {
	return me.httpClient.Do(method, me.basePath+path).Header("x-access-token", me.accessToken)
}
