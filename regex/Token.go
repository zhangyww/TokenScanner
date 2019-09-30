package regex

type  Token struct {
	Index			int
	Description		string
	LexerIndex		int
}

func NewToken(index int, description string, lexerIndex int) *Token {
	return &Token{
		Index:       index,
		Description: description,
		LexerIndex:  lexerIndex,
	}
}

func (this *Token) Equals(other *Token) bool {
	if other == nil {
		return false
	}
	return this.Index == other.Index
}
