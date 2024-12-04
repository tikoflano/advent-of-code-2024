package solution202403

import (
	"regexp"
	"strings"
)

type ParserState int

const (
	ParsingCommand ParserState = iota
	ParsingOperator
)

var validOperandRegex = regexp.MustCompile(`[0-9]`)

type parser struct {
	buffer           string
	Operations       []*Operation
	State            ParserState
	currentOperation *Operation
}

func NewParser() parser {
	return parser{
		State: ParsingCommand,
	}
}

func (parser *parser) NextChar(char string) {
	switch parser.State {
	case ParsingCommand:
		parser.parseCommand(char)
	case ParsingOperator:
		parser.parseOperator(char)
	}
}

func (parser *parser) parseCommand(char string) {
	if char != "(" {
		parser.buffer += char
		return
	}

	for _, command := range Commands {
		if strings.HasSuffix(parser.buffer, command.Keyword) {
			parser.currentOperation = &Operation{command: command}
			parser.State = ParsingOperator
			break
		}
	}

	parser.clearBuffer()
}

func (parser *parser) parseOperator(char string) {
	if char == "," || char == ")" {
		parser.currentOperation.addOperand(parser.buffer)
		parser.clearBuffer()

		if char == ")" {
			parser.commitOperation()
		}
	} else if validOperandRegex.MatchString(char) {
		parser.buffer += char
	} else {
		parser.dropOperation()
	}
}

func (parser *parser) clearBuffer() {
	parser.buffer = ""
}

func (parser *parser) commitOperation() {
	parser.Operations = append(parser.Operations, parser.currentOperation)
	parser.currentOperation = nil
	parser.State = ParsingCommand
}

func (parser *parser) dropOperation() {
	parser.currentOperation = nil
	parser.clearBuffer()
	parser.State = ParsingCommand
}
