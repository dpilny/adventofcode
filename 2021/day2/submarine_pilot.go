package day2

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Command struct {
	Operation string
	Argument  int
}

func parseCommands(path string) ([]Command, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(data), "\n")
	var commands []Command
	for _, rawLine := range raw {
		parts := strings.Split(rawLine, " ")
		argument, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		cmd := Command{
			parts[0],
			argument,
		}
		commands = append(commands, cmd)
	}
	return commands, nil
}

func calcMultipliedPosition(commands []Command) int {
	hor := 0
	ver := 0
	for _, cmd := range commands {
		if cmd.Operation == "forward" {
			hor += cmd.Argument
		} else if cmd.Operation == "down" {
			ver += cmd.Argument
		} else if cmd.Operation == "up" {
			ver -= cmd.Argument
		}
	}
	return hor * ver
}

func calcMultipliedPositionAimed(commands []Command) int {
	hor := 0
	ver := 0
	aim := 0
	for _, cmd := range commands {
		if cmd.Operation == "forward" {
			hor += cmd.Argument
			if aim != 0 {
				ver += aim * cmd.Argument
			}
		} else if cmd.Operation == "down" {
			aim += cmd.Argument
		} else if cmd.Operation == "up" {
			aim -= cmd.Argument
		}
	}
	return hor * ver
}
