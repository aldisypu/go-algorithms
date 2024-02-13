package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54acc128329e634e9a000362

func TraverseTCPStates(events []string) string {
	state := "CLOSED"

	for _, event := range events {
		switch state {
		case "CLOSED":
			if event == "APP_PASSIVE_OPEN" {
				state = "LISTEN"
			} else if event == "APP_ACTIVE_OPEN" {
				state = "SYN_SENT"
			} else {
				state = "ERROR"
			}
		case "LISTEN":
			if event == "RCV_SYN" {
				state = "SYN_RCVD"
			} else if event == "APP_SEND" {
				state = "SYN_SENT"
			} else if event == "APP_CLOSE" {
				state = "CLOSED"
			} else {
				state = "ERROR"
			}
		case "SYN_SENT":
			if event == "RCV_SYN" {
				state = "SYN_RCVD"
			} else if event == "RCV_SYN_ACK" {
				state = "ESTABLISHED"
			} else if event == "APP_CLOSE" {
				state = "CLOSED"
			} else {
				state = "ERROR"
			}
		case "SYN_RCVD":
			if event == "APP_CLOSE" {
				state = "FIN_WAIT_1"
			} else if event == "RCV_ACK" {
				state = "ESTABLISHED"
			} else if event == "RCV_FIN" {
				state = "CLOSE_WAIT"
			} else {
				state = "ERROR"
			}
		case "ESTABLISHED":
			if event == "APP_CLOSE" {
				state = "FIN_WAIT_1"
			} else if event == "RCV_FIN" {
				state = "CLOSE_WAIT"
			} else {
				state = "ERROR"
			}
		case "FIN_WAIT_1":
			if event == "RCV_FIN" {
				state = "CLOSING"
			} else if event == "RCV_FIN_ACK" {
				state = "TIME_WAIT"
			} else if event == "RCV_ACK" {
				state = "FIN_WAIT_2"
			} else {
				state = "ERROR"
			}
		case "CLOSING":
			if event == "RCV_ACK" {
				state = "TIME_WAIT"
			} else {
				state = "ERROR"
			}
		case "FIN_WAIT_2":
			if event == "RCV_FIN" {
				state = "TIME_WAIT"
			} else {
				state = "ERROR"
			}
		case "TIME_WAIT":
			if event == "APP_TIMEOUT" {
				state = "CLOSED"
			} else {
				state = "ERROR"
			}
		case "CLOSE_WAIT":
			if event == "APP_CLOSE" {
				state = "LAST_ACK"
			} else {
				state = "ERROR"
			}
		case "LAST_ACK":
			if event == "RCV_ACK" {
				state = "CLOSED"
			} else {
				state = "ERROR"
			}
		case "ERROR":
			return "ERROR"
		}
	}

	return state
}

func TestTraverseTCPStates(t *testing.T) {
	assert.Equal(t, "CLOSE_WAIT", TraverseTCPStates([]string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN"}))
	assert.Equal(t, "ESTABLISHED", TraverseTCPStates([]string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK"}))
	assert.Equal(t, "LAST_ACK", TraverseTCPStates([]string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN", "APP_CLOSE"}))
	assert.Equal(t, "SYN_SENT", TraverseTCPStates([]string{"APP_ACTIVE_OPEN"}))
	assert.Equal(t, "ERROR", TraverseTCPStates([]string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK", "APP_CLOSE", "APP_SEND"}))
}
