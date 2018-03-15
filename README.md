## RPC Listener for Nano Node Blocks
go listener to rpc callbacks from a nano node

set up nano node in config.json to send rpc callbacks:

...
"callback_address":"CONTAINER.IP.ADD.RESS",
"callback_port": "10077",
"callback_target": "/callback",
...

on ubuntu ufw allow 10077, on both sides.
