package gnlp

import (
	"strings"
)

type Token struct {
	string text
	int begin
	int end
}

type TokenizedSentence struct {
	Token[] tokens
}

type TokenizedText struct {
	TokenizedSentence[] sentences
}

type Tokenizer interface {
	tokenize() TokenizedText
}

type WhiteSpaceTokenizer() struct {
	Tokenizer
}

func (w WhiteSpaceTokenizer) tokenize(text string) {
	tokens := strings.Fields(text)
	return tokens
}