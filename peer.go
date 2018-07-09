package senclient

import (
	"crypto/tls"
)

/*
	Representation of the subethanet 'peer' concept.

	A Peer contains identifying information, and a collection of connections.

	A Peer (usually) has 2 sub-collections of spare connections, 'listeners' and 'speakers'. Listeners are
	connections waiting on a command, and speakers are connections that are ready to send a command to a
	remote listener. They act analogously to servers and clients, but are not dependent on the underlying
	tcp connection's 'direction'.
*/

type peer struct {
	idCert     tls.Certificate // Represents the ID of a person. Used to trust the clientCert.
	clientCert tls.Certificate // The client-specific cert used to handle encrypting connections.

	connections []tcpConnection
	listeners   []tcpConnection
	speakers    []tcpConnection
}
