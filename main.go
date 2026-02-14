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
	for pos < inputLen  {
		char := tokens[pos]
		fmt.Println(string(char))
		switch char {
		case ':':
		if tokens[pos+1] == ':' {
				result = append(result, DOUBLE_COLON)
			} else {
				result = append(result, COLON)
			}
		}
		pos++
	}
}
