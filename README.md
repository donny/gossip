gossip
======

A WebSocket back-end for a chat app.

It's based on the code from Gary Burd: [http://gary.burd.info/go-websocket-chat](http://gary.burd.info/go-websocket-chat).

It's simplified by removing the `http.HandleFunc` that serves `index.html`. It acts only as a WebSocket back-end. The front-end for `gossip` can be found in [gossip-web](http://github.com/donny/gossip-web). The `gossip` code has been modified to accept JSON messages for future expansion.
