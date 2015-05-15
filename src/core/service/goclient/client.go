/**
 * Copyright 2014 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package goclient

import (
	"fmt"
	"github.com/atnet/gof"
	"github.com/atnet/gof/net/jsv"
	"os"
)

var (
	_conn    *jsv.TCPConn
	Member   *memberClient
	Partner  *partnerClient
	Redirect *redirectClient
)

func Configure(net, addr string, c gof.App) {
	var err error
	_conn, err = jsv.Dial(net, addr)

	if err != nil {
		fmt.Println("[TCP]: Connect Refused,", addr)
		os.Exit(1)
	}

	jsv.Configure(c)
	Member = &memberClient{conn: _conn}
	Partner = &partnerClient{conn: _conn}
	Redirect = &redirectClient{conn: _conn}
}
