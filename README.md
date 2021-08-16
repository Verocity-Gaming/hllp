# hllp
Hell Let Loose RCON REST Server

This program serves as example usage for the [rcon](github.com/verocity-gaming/rcon) client.

`go build && ./hllp -addr <addr> -pass <pass> -port 12345`

# GET Requests

`GET /name` -> `Verocity Gaming // US East // verocity.gg`

# POST Requests

`POST / '{"cmd": "kick", "player": "xXplayerXx", "reason": "hacking"}' -> 200`