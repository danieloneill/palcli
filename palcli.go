package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	BaseURL  string
	Username string
	Password string
	Client   *http.Client
}

func NewClient(base, user, pass string) *Client {
	return &Client{
		BaseURL:  strings.TrimRight(base, "/"),
		Username: user,
		Password: pass,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) do(method, path string, body any) ([]byte, error) {
	var reader io.Reader

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, reader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(data))
	}

	return data, nil
}

func main() {
	base := flag.String("url", "http://127.0.0.1:8212/v1/api", "API base URL")
	user := flag.String("user", "admin", "username")
	pass := flag.String("pass", "", "password")
	flag.Parse()

	if *pass == "" {
		fmt.Println("password required")
		os.Exit(1)
	}

	if flag.NArg() < 1 {
		usage()
		return
	}

	cmd := flag.Arg(0)
	args := flag.Args()[1:]

	c := NewClient(*base, *user, *pass)

	switch cmd {

	case "ban":
		if len(args) < 2 {
			fmt.Println("ban <userid> <message>")
			return
		}
		msg := "Bye"
		if len(args) >= 2 {
			msg = strings.Join(args[1:], " ")
		}
		out(c.do("POST", "/ban", map[string]any{
			"userid": args[0],
			"message": msg,
		}))

	case "broadcast":
		if len(args) < 1 {
			fmt.Println("broadcast <message>")
			return
		}
		msg := strings.Join(args, " ")
		out(c.do("POST", "/broadcast", map[string]any{
			"message": msg,
		}))

	case "info":
		out(c.do("GET", "/info", nil))

	case "kick":
		if len(args) < 1 {
			fmt.Println("kick <userid> [reason]")
			return
		}
		msg := "Bye"
		if len(args) >= 2 {
			msg = strings.Join(args[1:], " ")
		}
		out(c.do("POST", "/kick", map[string]any{
			"userid": args[0],
            "message": msg,
		}))

	case "metrics":
		out(c.do("GET", "/metrics", nil))

	case "players":
		out(c.do("GET", "/players", nil))

	case "save":
		out(c.do("POST", "/save", nil))

	case "settings":
		out(c.do("GET", "/settings", nil))

	case "shutdown":
		delay := 10
		msg := "Server shutting down"
		if len(args) >= 1 {
			fmt.Sscanf(args[0], "%d", &delay)
		}
		if len(args) >= 2 {
			msg = strings.Join(args[1:], " ")
		}
		out(c.do("POST", "/shutdown", map[string]any{
			"waittime": delay,
			"message":  msg,
		}))

	case "stop":
		out(c.do("POST", "/stop", nil))

	case "unban":
		if len(args) < 1 {
			fmt.Println("unban <steamid>")
			return
		}
		out(c.do("POST", "/unban", map[string]any{
			"userid": args[0],
		}))


	default:
		usage()
	}
}

func out(data []byte, err error) {
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var pretty bytes.Buffer
	if json.Indent(&pretty, data, "", "  ") == nil {
		fmt.Println(pretty.String())
	} else {
		fmt.Println(string(data))
	}
}

func usage() {
	fmt.Println(`palcli [flags] <command>

Commands:
  ban <steamid> [reason]
  broadcast <message>
  info
  kick <userid> [reason]
  metrics
  players
  save
  settings
  shutdown [seconds] [message]
  stop
  unban <userid>

Flags:
  -url   API base (default http://127.0.0.1:8212/v1/api)
  -user  username
  -pass  password`)
}
