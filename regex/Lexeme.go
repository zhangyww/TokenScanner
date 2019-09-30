package regex

type Lexeme struct {
	TokenIndex		int
	Value			string
}

func NewLexeme(tokenIndex int, value string) *Lexeme{
	return &Lexeme{
		TokenIndex: tokenIndex,
		Value:      value,
	}
}
