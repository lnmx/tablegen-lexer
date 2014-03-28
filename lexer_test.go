package main

import (
	"io/ioutil"
	"testing"
)

var input = "tblgen_input.td"
var input_tokens = 139167

func TestLexer(t *testing.T) {
	buf, err := ioutil.ReadFile(input)

	if err != nil {
		t.Fatal(err)
	}

	toks := testParse(buf)

	if len(toks) != input_tokens {
		t.Fatal("expected", input_tokens, "tokens, got", len(toks))
	}
}

func BenchmarkLexer(b *testing.B) {

	buf, err := ioutil.ReadFile(input)

	if err != nil {
		b.Fatal(err)
	}

	// reset the bench timer to ignore file read time
	//
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		toks := testParse(buf)

		if len(toks) != input_tokens {
			b.Fatal("expected", input_tokens, "tokens, got", len(toks))
		}
	}
}

func testParse(buf []byte) (toks []Token) {

	toks = make([]Token, 0)

	lex := NewLexer(buf)

	for {
		tok := lex.NextToken()

		toks = append(toks, tok)

		if tok.Name == EOF {
			break
		}
	}

	return toks
}
