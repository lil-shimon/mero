package main

import "fmt"

const (
	DOUBLE_COLON = iota
	COLON
	DOUBLE_ARROW
	ARROW
	MINUS
	RETURN_ARROW
	AT
	PLUS
	ASTERISK
	SLASH
	COMMA
	LBRACKET
	RBRACKET
	LBRACE
	RBRACE
	SPACE
	IDENTIFIER
	DOT
	EQUAL
	STRING
	NUMBER
)

type Token struct {
	Kind int
	Value string
}

// TODO: IDENTIFIERの場合、なんなのかを識別するロジック

func main() {
	inputs := ":: name > 'mero'"
	analyze(inputs)
}

func analyze(tokens string) {
	inputLen := len(tokens)
	fmt.Println(inputLen)

	pos := 0
	var result []int
	for pos < inputLen {
		char := tokens[pos]
		fmt.Println(string(char))
		switch char {
		case ':':
			if pos+1 < inputLen && tokens[pos+1] == ':' {
				result = append(result, DOUBLE_COLON)
				pos++
			} else {
				result = append(result, COLON)
			}
		case '>':
			if pos+1 < inputLen && tokens[pos+1] == '>' {
				result = append(result, DOUBLE_ARROW)
				pos++
			} else {
				result = append(result, ARROW)
			}
		case '-':
			if pos+1 < inputLen && tokens[pos+1] == '>' {
				result = append(result, RETURN_ARROW)
				pos++
			} else {
				result = append(result, MINUS)
			}
		case '.':
			result = append(result, DOT)
		case '+':
			result = append(result, PLUS)
		case '*':
			result = append(result, ASTERISK)
		case '/':
			result = append(result, SLASH)
		case ',':
			result = append(result, COMMA)
		case '[':
			result = append(result, LBRACKET)
		case ']':
			result = append(result, RBRACKET)
		case '{':
			result = append(result, LBRACE)
		case '}':
			result = append(result, RBRACE)
		case ' ':
			result = append(result, SPACE)
		case '=':
			result = append(result, EQUAL)
		}
		pos++
	}
}
