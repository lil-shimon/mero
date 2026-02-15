package main

import (
	"testing"
)

// Lexer Tests

func TestAnalyze_VariableDeclaration(t *testing.T) {
	input := ":: name > 'mero'"
	tokens := analyze(input)

	expected := []Token{
		{Kind: DOUBLE_COLON, Value: "::"},
		{Kind: IDENTIFIER, Value: "name"},
		{Kind: ARROW, Value: ">"},
		{Kind: STRING, Value: "mero"},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, tok := range tokens {
		if tok.Kind != expected[i].Kind {
			t.Errorf("Token %d: expected Kind %d, got %d", i, expected[i].Kind, tok.Kind)
		}
		if tok.Value != expected[i].Value {
			t.Errorf("Token %d: expected Value %q, got %q", i, expected[i].Value, tok.Value)
		}
	}
}

func TestAnalyze_FunctionSignature(t *testing.T) {
	input := "fn sum > x : number >"
	tokens := analyze(input)

	expected := []Token{
		{Kind: FN, Value: "fn"},
		{Kind: IDENTIFIER, Value: "sum"},
		{Kind: ARROW, Value: ">"},
		{Kind: IDENTIFIER, Value: "x"},
		{Kind: COLON, Value: ":"},
		{Kind: IDENTIFIER, Value: "number"},
		{Kind: ARROW, Value: ">"},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, tok := range tokens {
		if tok.Kind != expected[i].Kind {
			t.Errorf("Token %d: expected Kind %d, got %d", i, expected[i].Kind, tok.Kind)
		}
		if tok.Value != expected[i].Value {
			t.Errorf("Token %d: expected Value %q, got %q", i, expected[i].Value, tok.Value)
		}
	}
}

func TestAnalyze_DoubleArrow(t *testing.T) {
	input := ">>"
	tokens := analyze(input)

	if len(tokens) != 1 {
		t.Fatalf("Expected 1 token (DOUBLE_ARROW), got %d", len(tokens))
	}

	if tokens[0].Kind != DOUBLE_ARROW {
		t.Errorf("Expected DOUBLE_ARROW (%d), got %d", DOUBLE_ARROW, tokens[0].Kind)
	}

	if tokens[0].Value != ">>" {
		t.Errorf("Expected Value '>>', got %q", tokens[0].Value)
	}
}

func TestAnalyze_ReturnArrow(t *testing.T) {
	input := "->"
	tokens := analyze(input)

	if len(tokens) != 1 {
		t.Fatalf("Expected 1 token (RETURN_ARROW), got %d", len(tokens))
	}

	if tokens[0].Kind != RETURN_ARROW {
		t.Errorf("Expected RETURN_ARROW (%d), got %d", RETURN_ARROW, tokens[0].Kind)
	}

	if tokens[0].Value != "->" {
		t.Errorf("Expected Value '->', got %q", tokens[0].Value)
	}
}

func TestAnalyze_NumberLiteral(t *testing.T) {
	input := ":: a > 123"
	tokens := analyze(input)

	expected := []Token{
		{Kind: DOUBLE_COLON, Value: "::"},
		{Kind: IDENTIFIER, Value: "a"},
		{Kind: ARROW, Value: ">"},
		{Kind: NUMBER, Value: "123"},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d", len(expected), len(tokens))
	}

	for i, tok := range tokens {
		if tok.Kind != expected[i].Kind {
			t.Errorf("Token %d: expected Kind %d, got %d", i, expected[i].Kind, tok.Kind)
		}
		if tok.Value != expected[i].Value {
			t.Errorf("Token %d: expected Value %q, got %q", i, expected[i].Value, tok.Value)
		}
	}
}

func TestAnalyze_SingleCharacterTokens(t *testing.T) {
	tests := []struct {
		input    string
		expected Token
	}{
		{"+", Token{Kind: PLUS, Value: "+"}},
		{"-", Token{Kind: MINUS, Value: "-"}},
		{"*", Token{Kind: ASTERISK, Value: "*"}},
		{"/", Token{Kind: SLASH, Value: "/"}},
		{",", Token{Kind: COMMA, Value: ","}},
		{".", Token{Kind: DOT, Value: "."}},
		{"[", Token{Kind: LBRACKET, Value: "["}},
		{"]", Token{Kind: RBRACKET, Value: "]"}},
		{"{", Token{Kind: LBRACE, Value: "{"}},
		{"}", Token{Kind: RBRACE, Value: "}"}},
	}

	for _, tt := range tests {
		tokens := analyze(tt.input)

		if len(tokens) != 1 {
			t.Fatalf("Input %q: expected 1 token, got %d", tt.input, len(tokens))
		}

		if tokens[0].Kind != tt.expected.Kind {
			t.Errorf("Input %q: expected Kind %d, got %d", tt.input, tt.expected.Kind, tokens[0].Kind)
		}

		if tokens[0].Value != tt.expected.Value {
			t.Errorf("Input %q: expected Value %q, got %q", tt.input, tt.expected.Value, tokens[0].Value)
		}
	}
}

// Parser Tests

func TestParse_VariableDeclaration(t *testing.T) {
	tokens := []Token{
		{Kind: DOUBLE_COLON, Value: "::"},
		{Kind: IDENTIFIER, Value: "name"},
		{Kind: ARROW, Value: ">"},
		{Kind: STRING, Value: "mero"},
	}

	nodes := parse(tokens)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got %d", len(nodes))
	}

	node := nodes[0]

	if node.Kind != NODE_VARIABLE_DECLARATION {
		t.Errorf("Expected Kind %q, got %q", NODE_VARIABLE_DECLARATION, node.Kind)
	}

	if node.Name != "name" {
		t.Errorf("Expected Name %q, got %q", "name", node.Name)
	}

	if node.Value != "mero" {
		t.Errorf("Expected Value %q, got %q", "mero", node.Value)
	}
}

func TestParse_FunctionDeclaration_SingleParam(t *testing.T) {
	tokens := []Token{
		{Kind: FN, Value: "fn"},
		{Kind: IDENTIFIER, Value: "sum"},
		{Kind: ARROW, Value: ">"},
		{Kind: IDENTIFIER, Value: "x"},
		{Kind: COLON, Value: ":"},
		{Kind: IDENTIFIER, Value: "number"},
		{Kind: ARROW, Value: ">"},
	}

	nodes := parse(tokens)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got %d", len(nodes))
	}

	node := nodes[0]

	if node.Kind != NODE_FUNCTION_DECLARATION {
		t.Errorf("Expected Kind %q, got %q", NODE_FUNCTION_DECLARATION, node.Kind)
	}

	if node.Name != "sum" {
		t.Errorf("Expected Name %q, got %q", "sum", node.Name)
	}

	if len(node.Params) != 1 {
		t.Fatalf("Expected 1 parameter, got %d", len(node.Params))
	}

	if node.Params[0].Name != "x" {
		t.Errorf("Expected Param[0].Name %q, got %q", "x", node.Params[0].Name)
	}

	if node.Params[0].Value != "number" {
		t.Errorf("Expected Param[0].Value %q, got %q", "number", node.Params[0].Value)
	}
}

func TestParse_FunctionDeclaration_MultipleParams(t *testing.T) {
	tokens := []Token{
		{Kind: FN, Value: "fn"},
		{Kind: IDENTIFIER, Value: "sum"},
		{Kind: ARROW, Value: ">"},
		{Kind: IDENTIFIER, Value: "x"},
		{Kind: COLON, Value: ":"},
		{Kind: IDENTIFIER, Value: "number"},
		{Kind: COMMA, Value: ","},
		{Kind: IDENTIFIER, Value: "y"},
		{Kind: COLON, Value: ":"},
		{Kind: IDENTIFIER, Value: "number"},
		{Kind: ARROW, Value: ">"},
	}

	nodes := parse(tokens)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got %d", len(nodes))
	}

	node := nodes[0]

	if node.Kind != NODE_FUNCTION_DECLARATION {
		t.Errorf("Expected Kind %q, got %q", NODE_FUNCTION_DECLARATION, node.Kind)
	}

	if node.Name != "sum" {
		t.Errorf("Expected Name %q, got %q", "sum", node.Name)
	}

	if len(node.Params) != 2 {
		t.Fatalf("Expected 2 parameters, got %d", len(node.Params))
	}

	if node.Params[0].Name != "x" {
		t.Errorf("Expected Param[0].Name %q, got %q", "x", node.Params[0].Name)
	}

	if node.Params[0].Value != "number" {
		t.Errorf("Expected Param[0].Value %q, got %q", "number", node.Params[0].Value)
	}

	if node.Params[1].Name != "y" {
		t.Errorf("Expected Param[1].Name %q, got %q", "y", node.Params[1].Name)
	}

	if node.Params[1].Value != "number" {
		t.Errorf("Expected Param[1].Value %q, got %q", "number", node.Params[1].Value)
	}
}

func TestParse_FunctionDeclaration_WithReturnArrow(t *testing.T) {
	// fn double > x : number >
	//     -> x
	// >
	tokens := []Token{
		{Kind: FN, Value: "fn"},
		{Kind: IDENTIFIER, Value: "double"},
		{Kind: ARROW, Value: ">"},
		{Kind: IDENTIFIER, Value: "x"},
		{Kind: COLON, Value: ":"},
		{Kind: IDENTIFIER, Value: "number"},
		{Kind: ARROW, Value: ">"},
		{Kind: RETURN_ARROW, Value: "->"},
		{Kind: IDENTIFIER, Value: "x"},
		{Kind: ARROW, Value: ">"},
	}

	nodes := parse(tokens)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got %d", len(nodes))
	}

	node := nodes[0]

	if node.Kind != NODE_FUNCTION_DECLARATION {
		t.Errorf("Expected Kind %q, got %q", NODE_FUNCTION_DECLARATION, node.Kind)
	}

	if node.Name != "double" {
		t.Errorf("Expected Name %q, got %q", "double", node.Name)
	}

	if len(node.Params) != 1 {
		t.Fatalf("Expected 1 parameter, got %d", len(node.Params))
	}

	if node.Params[0].Name != "x" {
		t.Errorf("Expected Param[0].Name %q, got %q", "x", node.Params[0].Name)
	}

	if node.Params[0].Value != "number" {
		t.Errorf("Expected Param[0].Value %q, got %q", "number", node.Params[0].Value)
	}

	if len(node.Body) != 1 {
		t.Fatalf("Expected 1 body node, got %d", len(node.Body))
	}

	if node.Body[0].Kind != NODE_RETURN_STATEMENT {
		t.Errorf("Expected Body[0].Kind %q, got %q", NODE_RETURN_STATEMENT, node.Body[0].Kind)
	}

	if node.Body[0].Value != "x" {
		t.Errorf("Expected Body[0].Value %q, got %q", "x", node.Body[0].Value)
	}
}

// Lexer: print@name のトークン化
func TestAnalyze_PrintStatement(t *testing.T) {
	input := "print@name"
	tokens := analyze(input)

	expected := []Token{
		{Kind: PRINT, Value: "print"},
		{Kind: AT, Value: "@"},
		{Kind: IDENTIFIER, Value: "name"},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d: %v", len(expected), len(tokens), tokens)
	}

	for i, tok := range tokens {
		if tok.Kind != expected[i].Kind {
			t.Errorf("Token %d: expected Kind %d, got %d", i, expected[i].Kind, tok.Kind)
		}
		if tok.Value != expected[i].Value {
			t.Errorf("Token %d: expected Value %q, got %q", i, expected[i].Value, tok.Value)
		}
	}
}

// Parser: print@name のパース
func TestParse_PrintStatement(t *testing.T) {
	tokens := []Token{
		{Kind: PRINT, Value: "print"},
		{Kind: AT, Value: "@"},
		{Kind: IDENTIFIER, Value: "name"},
	}

	nodes := parse(tokens)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got %d", len(nodes))
	}

	node := nodes[0]

	if node.Kind != NODE_PRINT {
		t.Errorf("Expected Kind %q, got %q", NODE_PRINT, node.Kind)
	}

	if node.Value != "name" {
		t.Errorf("Expected Value %q, got %q", "name", node.Value)
	}
}

// Lexer: 変数宣言 + print の複合
func TestAnalyze_VariableAndPrint(t *testing.T) {
	input := ":: name > 'mero'\nprint@name"
	tokens := analyze(input)

	expected := []Token{
		{Kind: DOUBLE_COLON, Value: "::"},
		{Kind: IDENTIFIER, Value: "name"},
		{Kind: ARROW, Value: ">"},
		{Kind: STRING, Value: "mero"},
		{Kind: PRINT, Value: "print"},
		{Kind: AT, Value: "@"},
		{Kind: IDENTIFIER, Value: "name"},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens, got %d: %v", len(expected), len(tokens), tokens)
	}

	for i, tok := range tokens {
		if tok.Kind != expected[i].Kind {
			t.Errorf("Token %d: expected Kind %d, got %d", i, expected[i].Kind, tok.Kind)
		}
		if tok.Value != expected[i].Value {
			t.Errorf("Token %d: expected Value %q, got %q", i, expected[i].Value, tok.Value)
		}
	}
}

// Evaluator Tests

func TestEval_VariableDeclaration(t *testing.T) {
	nodes := []Node{
		{Kind: NODE_VARIABLE_DECLARATION, Name: "name", Value: "mero"},
	}
	env := map[string]string{}
	eval(nodes, env)

	if env["name"] != "mero" {
		t.Errorf("Expected env[\"name\"] = %q, got %q", "mero", env["name"])
	}
}

func TestEval_MultipleVariables(t *testing.T) {
	nodes := []Node{
		{Kind: NODE_VARIABLE_DECLARATION, Name: "a", Value: "hello"},
		{Kind: NODE_VARIABLE_DECLARATION, Name: "b", Value: "world"},
	}
	env := map[string]string{}
	eval(nodes, env)

	if env["a"] != "hello" {
		t.Errorf("Expected env[\"a\"] = %q, got %q", "hello", env["a"])
	}
	if env["b"] != "world" {
		t.Errorf("Expected env[\"b\"] = %q, got %q", "world", env["b"])
	}
}

func TestAnalyze_IdentifierWithUnderscore(t *testing.T) {
	input := "number_one"
	tokens := analyze(input)

	if len(tokens) != 1 {
		t.Fatalf("Expected 1 token, got %d: %v", len(tokens), tokens)
	}

	if tokens[0].Kind != IDENTIFIER {
		t.Errorf("Expected Kind IDENTIFIER (%d), got %d", IDENTIFIER, tokens[0].Kind)
	}

	if tokens[0].Value != "number_one" {
		t.Errorf("Expected Value %q, got %q", "number_one", tokens[0].Value)
	}
}

func TestAnalyze_EqualEqual(t *testing.T) {
	input := "=="
	tokens := analyze(input)

	// Currently == produces two EQUAL tokens
	// This test documents the current behavior
	// TODO: may need DOUBLE_EQUAL token for comparison
	if len(tokens) != 2 {
		t.Fatalf("Expected 2 tokens, got %d: %v", len(tokens), tokens)
	}

	if tokens[0].Kind != EQUAL || tokens[1].Kind != EQUAL {
		t.Errorf("Expected two EQUAL tokens, got %v", tokens)
	}
}

func TestParse_VariableDeclarationWithNumber(t *testing.T) {
	tokens := []Token{
		{Kind: DOUBLE_COLON, Value: "::"},
		{Kind: IDENTIFIER, Value: "a"},
		{Kind: ARROW, Value: ">"},
		{Kind: NUMBER, Value: "123"},
	}

	nodes := parse(tokens)

	if len(nodes) != 1 {
		t.Fatalf("Expected 1 node, got %d", len(nodes))
	}

	node := nodes[0]

	if node.Kind != NODE_VARIABLE_DECLARATION {
		t.Errorf("Expected Kind %q, got %q", NODE_VARIABLE_DECLARATION, node.Kind)
	}

	if node.Name != "a" {
		t.Errorf("Expected Name %q, got %q", "a", node.Name)
	}

	if node.Value != "123" {
		t.Errorf("Expected Value %q, got %q", "123", node.Value)
	}
}
