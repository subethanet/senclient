This document outlines the SEN protocol. It is the **canonical source of truth**, and any discrepancies (including the code) should be considered in need of updating.

All commands start with

`SEN <major>`

where `<major>` is the protocol version. Expect frequent backwards incompatibility with version 0.

# v0

All commands in v0 start with

`SEN 0 <sender>`

Assume that v0 commands have that prefix. Almost all commands start with

`SEN 0 <sender> <destination>`

but `<destination>` will be explicitly noted when present.

## Whois

`WHOIS`

`WHOIS` returns an `IAM`.

## Iam

`IAM <id>`

## Peer

`<destination> PEER [target <self address> [NEAR, LEFT, RIGHT]]`

A PEER request attempts to peer with another client. The command may simply be

`<bob's id> PEER`

when attempting to peer with Bob after opening a connection.

If the prospective peer is not connected, a message is carried via an existing peer.

`<existing peer> PEER <target> <self address> [NEAR, LEFT, RIGHT]`

`NEAR`, `LEFT`, and `RIGHT` are optional parameters that allow non-exact matches. `LEFT` seeks the closest possible peer to the "left" of the target, similarly with `RIGHT`. `NEAR` seeks the closest possible peer to the target, regardless of the direction.

## Data

`DATA <something>`

`DATA` sends arbitrary data. It can be used to construct higher-level protocols on top of SEN.