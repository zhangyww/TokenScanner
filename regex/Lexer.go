package regex

type Lexer struct {
	Lexicon			*Lexicon
	BaseLexer		*Lexer
	TokenInfos		[]*TokenInfo
	Index			int
	Level			int
	Children		[]*Lexer
}

func NewLexer(lexicon *Lexicon, index int) *Lexer {
	return NewLexerFromBase(lexicon, index, nil)
}

func NewLexerFromBase(lexicon *Lexicon, index int, baseLexer *Lexer) *Lexer {
	this := &Lexer{
		Lexicon:    lexicon,
		BaseLexer:  baseLexer,
		TokenInfos: []*TokenInfo{},
		Index:      index,
		Level:      0,
		Children:   []*Lexer{},
	}

	if baseLexer != nil {
		this.Level = baseLexer.Level + 1
		baseLexer.Children = append(baseLexer.Children, this)
	}

	return this
}

func (this *Lexer) DefineTokenWithDescription(exp IRegex, description string) *Token {
	indexInState := len(this.TokenInfos)

	tokenInfo := this.Lexicon.AddToken(exp, this, indexInState, description)

	this.TokenInfos = append(this.TokenInfos, tokenInfo)

	return tokenInfo.Token
}

func (this *Lexer) DefineToken(exp IRegex) *Token {
	return this.DefineTokenWithDescription(exp, "null")
}

func (this *Lexer) CreateSubLexer() *Lexer {
	return this.Lexicon.DefineLexer(this)
}