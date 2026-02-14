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
	for pos < inputLen  {
		char := tokens[pos]
		fmt.Println(string(char))
		pos++
	}
}
