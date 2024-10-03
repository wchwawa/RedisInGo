package main

import (
	"fmt"
	"strconv"
	"strings"
)

var CommandHandlers = map[string]func([]string) ([]byte, error){
	"ping": handlePING,
	"echo": handleECHO,
}

func validCommand(command string) bool {
	_, ok := CommandHandlers[strings.ToLower(command)]
	return ok
}
func parseResponseCommand(command []string) ([]byte, error) {
	if len(command) == 0 {
		return nil, fmt.Errorf("empty command")
	}
	CommandType := strings.ToLower(command[0])
	fmt.Println("====================" + CommandType + "====================")
	if validCommand(CommandType) == false {
		return nil, fmt.Errorf("invalid command")
	}

	return CommandHandlers[CommandType](command)
}

func handlePING(command []string) ([]byte, error) {
	return []byte("+PONG\r\n"), nil
}

func handleECHO(command []string) ([]byte, error) {
	if len(command) != 2 {
		return nil, fmt.Errorf("missing argument")
	}
	sb := strings.Builder{}
	sb.WriteString("$")
	length := len(command[1])
	sb.WriteString(strconv.Itoa(length))
	sb.WriteString("\r\n")
	sb.WriteString(command[1])
	sb.WriteString("\r\n")
	return []byte(sb.String()), nil
}
