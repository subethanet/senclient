package peer

import (
	"senclient/con"
)

/*
	Representation of the subethanet 'peer' concept.

	A Peer contains identifying information, and a collection of connections.

	A Peer (usually) has 2 sub-collections of spare connections, 'listeners' and 'speakers'. Listeners are
	connections waiting on a command, and speakers are connections that are ready to send a command to a
	remote listener. They act analogously to servers and clients, but are not dependent on the underlying
	tcp connection's direction.
*/


type Peer struct {
	listeners []con.Connection
	speakers []con.Connection
	// Should also track all peered Connections via the peer for a graceful de-peering.
}