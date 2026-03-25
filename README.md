# palcli
A simple CLI for Palworld servers via REST API written in Go

--

## Usage

Full options are:
```
palcli [flags] <command>

Commands:
  ban <steamid>
  broadcast <message>
  info
  kick <steamid>
  metrics
  players
  save
  shutdown [seconds] [message]

Flags:
  -url   API base (default http://127.0.0.1:8212/v1/api)
  -user  username
  -pass  password
```

The only required flag is **-pass**

Basic usage is typically as follows:

```
palcli -pass=secret info
palcli -pass=secret metrics
palcli -pass=secret players
palcli -pass=secret kick 7656119XXXXXXXXXX
palcli -pass=secret broadcast "server restart in 5 minasauruses"
palcli -pass=secret save
palcli -pass=secret shutdown 60 "latte nerds"
```

If your server is not on localhost, pass `-url=http://1.2.3.4:8212/v1/api` as appropriate.

## Building

`go build -o palcli`

## Automation

As an example, if you place palcli in */usr/local/bin/* and wish to automate polite/safe shutdowns, you could add this to the **[Service]** section of your *palworld.service* (or whatever) file:

```
# Give enough time for the warning + shutdown
KillSignal=SIGINT
KillMode=mixed
TimeoutStopSec=120

ExecStop=/usr/local/bin/palcli -pass=secret broadcast "Server shutting down in 60 seconds"
ExecStop=/bin/sleep 60
ExecStop=/usr/local/bin/palcli -pass=secret save
ExecStop=/usr/local/bin/palcli -pass=secret shutdown 5 "Shutting down now"
```

