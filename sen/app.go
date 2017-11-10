package sen

import "manifold/sen/con/tcpserver"

/*
	Go library for the subethanet protocol.
*/


func Start() tcpserver.Server {
	server := tcpserver.Create(4242)
	server.Start()
	return server
}