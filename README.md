# simple-ledger
A ledger I am building for fun/example of my skills

### Updates

This is a work-in-progress.

This chain will be permissioned, distributed, and will support sidechains somehow.

there is a http server running on port 6789

### Running the daemon

```
go build && ./simple-ledger <chainName>
```

### HTTP API

GET /info

GET /block/<blockHeight>

GET /create/keypair

POST /permission/grant
```
{
	"address": "0350af209a983d4157e836b00be315624f849aaad950ca9b055597fdc97aa7edb7",
	"actions": ["send", "receive", "mine"],
}
```
