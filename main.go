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
	Kind  int
	Value string
}

// TODO: IDENTIFIERの場合、なんなのかを識別するロジック

func main() {
	inputs := ":: name > 'mero'"
	analyze(inputs)
}

// Lexer
func analyze(tokens string) {
	inputLen := len(tokens)
	pos := 0
	var result []Token

	for pos < inputLen {
		char := tokens[pos]
		switch char {
		case ':':
			if pos+1 < inputLen && tokens[pos+1] == ':' {
				result = append(result, Token{Kind: DOUBLE_COLON, Value: "::"})
				pos++
			} else {
				result = append(result, Token{Kind: COLON, Value: ":"})
			}
		case '>':
			if pos+1 < inputLen && tokens[pos+1] == '>' {
				result = append(result, Token{Kind: DOUBLE_ARROW, Value: ">>"})
				pos++
			} else {
				result = append(result, Token{Kind: ARROW, Value: ">"})
			}
		case '-':
			if pos+1 < inputLen && tokens[pos+1] == '>' {
				result = append(result, Token{Kind: RETURN_ARROW, Value: "->"})
				pos++
			} else {
				result = append(result, Token{Kind: MINUS, Value: "-"})
			}
		case '.':
			result = append(result, Token{Kind: DOT, Value: "."})
		case '+':
			result = append(result, Token{Kind: PLUS, Value: "+"})
		case '*':
			result = append(result, Token{Kind: ASTERISK, Value: "*"})
		case '/':
			result = append(result, Token{Kind: SLASH, Value: "/"})
		case ',':
			result = append(result, Token{Kind: COMMA, Value: ","})
		case '[':
			result = append(result, Token{Kind: LBRACKET, Value: "["})
		case ']':
			result = append(result, Token{Kind: RBRACKET, Value: "]"})
		case '{':
			result = append(result, Token{Kind: LBRACE, Value: "{"})
		case '}':
			result = append(result, Token{Kind: RBRACE, Value: "}"})
		case ' ':
			result = append(result, Token{Kind: SPACE, Value: " "})
		case '=':
			result = append(result, Token{Kind: EQUAL, Value: "="})
		default:
			if char == '\'' {
				// 開始の'をスキップ
				pos++
				start := pos

				for pos < inputLen && tokens[pos] != '\'' {
					pos++
				}
				result = append(result, Token{Kind: STRING, Value: tokens[start:pos]})
			} else if char >= '0' && char <= '9' {
				start := pos

				for pos < inputLen && tokens[pos] >= '0' && tokens[pos] <= '9' {
					pos++
				}
				result = append(result, Token{Kind: NUMBER, Value: tokens[start:pos]})
				// 最後のループ分進んでいるので戻す
				pos--
			} else if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
				start := pos

				for pos < inputLen && tokens[pos] >= 'a' && tokens[pos] <= 'z' || tokens[pos] >= 'A' && tokens[pos] <= 'Z' {
					pos++
				}
				result = append(result, Token{Kind: IDENTIFIER, Value: tokens[start:pos]})
				pos--
			}
		}
		pos++
	}
	fmt.Println(result)
}
