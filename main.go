package main

import (
	"fmt"
	"os"
)

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
	IDENTIFIER
	DOT
	EQUAL
	STRING
	NUMBER
	FN
	IF
	ELSE
	TYPE
)

type Token struct {
	Kind  int
	Value string
}

// TODO: IDENTIFIERの場合、なんなのかを識別するロジック(e.g. if, type)

var keywords = map[string]int{
	"fn": FN,
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: mero <file.mero>")
		return
	}

	content, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	inputs := string(content)
	tokens := analyze(inputs)
	nodes := parse(tokens)

	env := map[string]string{}

	eval(nodes, env)
}

// Lexer
func analyze(tokens string) []Token {
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

				value := tokens[start:pos]

				if kind, ok := keywords[value]; ok {
					result = append(result, Token{Kind: kind, Value: value})
				} else {
					result = append(result, Token{Kind: IDENTIFIER, Value: tokens[start:pos]})
				}
				pos--
			}
		}
		pos++
	}
	fmt.Println(result)
	return result
}

// Parser

// AST Node
type Node struct {
	Kind   string
	Name   string
	Value  string
	Params []Node
	Body   []Node
}

const (
	NODE_VARIABLE_DECLARATION = "VariableDeclaration"
	NODE_FUNCTION_DECLARATION = "FunctionDeclaration"
	NODE_RETURN_STATEMENT = "ReturnStatement"
)

func parse(tokens []Token) []Node {
	var result []Node
	pos := 0

	for pos < len(tokens) {
		token := tokens[pos]
		switch token.Kind {
		case DOUBLE_COLON:
			pos++ // IDENTIFIER
			name := tokens[pos].Value
			pos++ // ARROW
			pos++ // STRING or NUMBER
			value := tokens[pos].Value
			result = append(result, Node{Kind: NODE_VARIABLE_DECLARATION, Name: name, Value: value})
		case FN:
			pos++ // IDENTIFIER
			name := tokens[pos].Value
			pos++ // ARROW
			pos++
			kind := tokens[pos].Kind
			switch kind {
			case IDENTIFIER:
				params := parseParams(tokens, &pos)

				var body [] Node

				if pos+1 < len(tokens) && tokens[pos+1].Kind == RETURN_ARROW {
					pos++ // ARROW
					pos++ // RETURN_ARROW

					body = append(body, Node{Kind: NODE_RETURN_STATEMENT, Value: tokens[pos].Value})
				}

				result = append(result, Node{Kind: NODE_FUNCTION_DECLARATION, Params: params, Name: name, Body: body})
			case RETURN_ARROW:
				var body [] Node

	 			pos++ // RETURN_ARROW
				for tokens[pos].Kind != ARROW {
					bodyValue := tokens[pos].Value
					body = append(body, Node{Value: bodyValue})
					pos++
				}
				
				result = append(result, Node{ Kind: NODE_RETURN_STATEMENT, Body: body})
			}
		}

		pos++
	}

	fmt.Println("parse result", result)
	return result
}

func parseParams(tokens []Token, pos *int) []Node {
	var params []Node

	for tokens[*pos].Kind != ARROW {
		paramName := tokens[*pos].Value
		*pos = *pos + 1 // COLON
		*pos = *pos + 1 // TYPE
		paramType := tokens[*pos].Value
		params = append(params, Node{Name: paramName, Value: paramType})

		*pos = *pos + 1

		if tokens[*pos].Kind == COMMA {
			*pos = *pos + 1 // SKIP COMMA
		}
	}

	return params
}

// Evaluate

func eval(nodes []Node, env map[string]string) {
	for _, node := range nodes {
		switch node.Kind {
		case NODE_VARIABLE_DECLARATION:
			env[node.Name] = node.Value
		}
	}
}
