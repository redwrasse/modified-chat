# modified-chat
Security modifications (wip) to the gorilla/websocket [chat example](https://github.com/gorilla/websocket/tree/master/examples/chat)

Follows the notes: https://www.redwrasse.io/notes/designnotesgochat


**Concepts**

A chat app requires

i) a client can join (requires a session with a server)

ii) a client can send a message (requires server to accept messages)

iii) server needs to know all existing clients

iv) client can send message to all other clients (requires iii), then client can send to server who can send to all other clients)

**Original chat architecture**

The original architecture makes uses of the websocket protocol.

Websocket enables two-way communication between a client and server. So a message from a client is sent down to the other clients by the server. See RFC 6455 for Websockets.

As expected each client has its own websocket connection with the server, which maintains a list of all clients and connections. Each client upon opening a connection registers a callback for messages from the server and appends them to a local log. The entity responsible for maintaining all client behaviors is called the Hub, and that for communicating with each client a Client (all on the server, of course).

One goroutine exists for the Hub, two for each Client (one for reading, the other for writing). Goroutines are golang’s unit of concurrency, channels are the means for them to communicate. See golang concurrency docs. Clearly the channels enable a message to go from a Client, through the Hub, and out to all remaining Clients. Perhaps looks something like below

```
client 2 --- Client 2          Client 4 --- client 4
                        \  /      
                        Hub
                        /  \    
client 1 --- Client 1          Client 3 --- client 3
```

**Security threats & mechanisms**

i) Client sends malware/undesired data in message.

ii) Enforce ‘groups’: subsets of clients.

iii) Create a form of role-based access control. Possible roles: those who can send binary data to other clients, those who can invite/expel other clients, those who can control/manipulate server properties.

iv) DDOS the server; => DDOS other clients?

v) Authentication layer in front, using a 3rd party. For example, Google sign-in.

vi) Message authentication between clients (message authentication scheme)

vii) Confidentiality betwee clients.
