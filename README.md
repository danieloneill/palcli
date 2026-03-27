# palcli
A simple CLI for Palworld servers via REST API written in Go

--

## Usage

Full options are:
```
palcli [flags] <command>

Commands:
  announce <message>
  ban <userid> [reason]
  info
  kick <userid> [reason]
  metrics
  players
  save
  shutdown [seconds] [message]
  stop
  unban <userid>

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
palcli -pass=secret kick 7656119XXXXXXXXXX "eat a big pile of mcdonald's, undesirable!"
palcli -pass=secret announce "server restart in 5 minasauruses"
palcli -pass=secret save
palcli -pass=secret shutdown 60 "latte nerds"
```

If your server is not on localhost, pass `-url=http://1.2.3.4:8212/v1/api` as appropriate.

## Building


If you haven't already cloned:
`go install github.com/danieloneill/palcli@latest`

Then find it in *~/go/bin/* (or wherever you place your quality Go binaries).

Otherwise:

```
$ git clone --depth 1 https://github.com/danieloneill/palcli.git
Cloning into 'palcli'...
remote: Enumerating objects: 8, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (7/7), done.
remote: Total 8 (delta 0), reused 3 (delta 0), pack-reused 0 (from 0)
Receiving objects: 100% (8/8), 14.99 KiB | 269.00 KiB/s, done.
$ cd palcli
$ go build -o palcli
```

... but I mean, if I have to explain how to cp it to a PATH dir ... ask Claude.

## Automation

As an example, if you place palcli in */usr/local/bin/* and wish to automate polite/safe shutdowns, you could add this to the **[Service]** section of your *palworld.service* (or whatever) file:

```
# Give enough time for the warning + shutdown
KillSignal=SIGINT
KillMode=mixed
TimeoutStopSec=120

ExecStop=/usr/local/bin/palcli -pass=secret announce "Server shutting down in 60 seconds"
ExecStop=/bin/sleep 60
ExecStop=/usr/local/bin/palcli -pass=secret save
ExecStop=/usr/local/bin/palcli -pass=secret shutdown 5 "Shutting down now"
```

As for Windows, absolutely nfi.
