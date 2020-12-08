package halter

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Argument  int
}

type Program struct {
	Code         []Instruction
	VisitedLines map[int]bool
	Acc          int
	CurrentLine  int
	LastLine     int
}

func parseInstruction(line string) (*Instruction, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid instruction - needs to consist of operation and argument")
	}

	op := parts[0]
	if op != "nop" && op != "acc" && op != "jmp" {
		return nil, errors.New("invalid instruction - unknown operation: " + op)
	}

	sign := string(parts[1][0])
	argument, err := strconv.ParseInt(parts[1][1:len(parts[1])], 10, 64)
	if err != nil {
		return nil, errors.New("invalid instruction - invalid argument format: " + err.Error())
	}
	if sign == "-" {
		argument *= -1
	} else if sign != "+" {
		return nil, errors.New("invalid instruction - invalid sign for argument: " + sign)
	}
	return &Instruction{
		Operation: op,
		Argument:  int(argument),
	}, nil
}

func Initialize(codePath string) (*Program, error) {
	data, err := ioutil.ReadFile(codePath)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(data), "\n")
	var lines []Instruction
	for _, rawLine := range raw {
		line, err := parseInstruction(rawLine)
		if err != nil {
			return nil, err
		}
		lines = append(lines, *line)
	}
	return &Program{
		Code:         lines,
		LastLine:     len(raw),
		VisitedLines: map[int]bool{},
	}, nil
}

func (p *Program) processNextLine() bool {
	if _, ok := p.VisitedLines[p.CurrentLine]; ok {
		log.Println("detected infinite loop - acc val is:", p.Acc)
		return true
	}
	p.VisitedLines[p.CurrentLine] = true
	ins := p.Code[p.CurrentLine]
	switch ins.Operation {
	case "jmp":
		p.CurrentLine += ins.Argument
		break
	case "acc":
		p.Acc += ins.Argument
		p.CurrentLine++
		break
	case "nop":
		p.CurrentLine++
		break
	}
	log.Println("processed line - values are - acc:", p.Acc, "next line is:", p.CurrentLine)
	return p.CurrentLine == p.LastLine
}

func (p *Program) run() {
	for {
		finished := p.processNextLine()
		if finished {
			log.Println("finished program or detected infinite loop")
			break
		}
	}
}
