# Terminology

* **Carrier**: a user who receives a message that they are not the destination for.
* **Client**: a running instance of senclient.
* **Connection**: a TCP connection between 2 peers. Connections can be **inbound** ("this" client is acting as the TCP server) or **outbound** (the other client is acting as the TCP server).
* **Destination**: the user that a message is for. Carriers are not destinations, even though they have a wrapped message addressed to them.
* **Peer**: a subethanet client with a network and application-level connection. A client can be connected but not peered (EG prior to the PEER handshake).

# Client lookup

The SEN client can't realistically be connected to the entire network at once (network size, CAP theorem, etc). The subethanet uses a [distributed hash table](https://en.wikipedia.org/wiki/Distributed_hash_table) of user IDs to find other users in the network.

Broadly, a distributed hash table is a scheme for self-organizing peered connections, such that from any client, there is a clear path "closer" to the target. It helps to visualize as a circle, where any range of IDs could be considered to loop around (EG integer overflow).

Clients maintain a "fan" of peers with various IDs, with a bias to IDs that are closer to their own. In particular, clients periodically try to find the IDs that are asymptotically closest to their own. For example, ID 10 would try to find the highest ID <10 and the lowest ID >10.

When a client receives a message as a carrier (IE it isn't the destination), it checks which one of its peers is "closer". The "asymptotically close" peers are critical here, as they provide a (weak) guarantee of structure in the network. Any client should be able to pass "left" or "right" without "overshooting" the target ID. 

# Communication lifecycle

Alice wants to send a message to Bob.

## Alice and Bob are peers

1. Alice sends a message to Bob, encrypted to his ID.

## Alice and Bob are not peers

1. Alice looks at her peers to find the peer with the ID "closest" to Bob's, which is Eve.
2. Alice sends a PEER <Bob's ID> request to Eve.
3. Eve receives the request.
    * If Eve is not paired with Bob, Eve repeats the process of step 2 with one of her peers.
    * If Eve is peered with Bob, delivers Alice's PEER message.
4. Bob receives Alive's PEER request. Bob attempts to connect directly to Alice's IP address and make a PEER request.
    * If this succeeds, Alice and Bob are now peers. (See "Alice and Bob are peers" section)
6. If Bob cannot connect to Alice (EG Alice doesn't have an exposed port), Bob sends an INFO message (including his IP address) to Alice.
7. Alice recieves Bob's INFO message, and attempts to connect directly to BOB and make a PEER request.
    * If this ALSO fails, see "Alice and Bob cannot peer".
    
## Alice and Bob cannot peer

Peering relies on a direct network connection (note: a "virtual peering" concept may be introduced at some point). If neither party can connect due to network issues, they can still communicate indirectly.

Alice can send messages to Bob, using Eve as the carrier, just as she can send meta/command messages. It is slower and risks the encrypted messages being siphoned.