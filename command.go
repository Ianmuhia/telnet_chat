package main

type commandID int

type command struct {
	id     commandID
	client *client
	args   []string
}

const (
	CMD_IAN commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)
