package regex

type Lexicon struct {
	TokenInfos		[]*TokenInfo
	Lexers			[]*Lexer
	DefaultLexer	*Lexer
}

func NewLexicon() *Lexicon {
	this := &Lexicon{
		TokenInfos:   []*TokenInfo{},
		Lexers:       []*Lexer{},
		DefaultLexer: nil,
	}

	this.DefaultLexer = NewLexer(this, 0)
	this.Lexers = append(this.Lexers, this.DefaultLexer)

	return this
}

func (this *Lexicon) AddToken(exp IRegex, lexer *Lexer, indexInState int, description string) *TokenInfo {
	index := len(this.TokenInfos)
	token := NewToken(index, description, lexer.Index)
	tokenInfo := NewTokenInfo(exp, this, lexer, token)

	this.TokenInfos = append(this.TokenInfos, tokenInfo)

	return tokenInfo
}

func (this *Lexicon) DefineLexer(baseLexer *Lexer) *Lexer {
	index := len(this.Lexers)
	newLexer := NewLexerFromBase(this, index, baseLexer)
	this.Lexers = append(this.Lexers, newLexer)

	return newLexer
}


func (this *Lexicon) CreateCompactCharsetManager() *CompactCharsetManager {
	tokenInfos := this.TokenInfos

	uncompactableCharSet := NewRuneSet()

	for idx, _ := range tokenInfos {
		uncompactableCharSet.Union(tokenInfos[idx].Exp.GetUncompactableCharset())
	}

	charClass := rune(1)
	compactClassTable := make([]rune, 65536)
	for c, _ := range uncompactableCharSet.Data {
		compactClassTable[c] = charClass
		charClass ++
	}

	return NewCompatCharsetManager(compactClassTable, charClass)
}

func (this *Lexicon) CreateScannerInfo() *ScannerInfo {
	dfaModel := CreateDFAModel(this)
	t := Compress(dfaModel)

	return NewSannerInfo(t.TransitionTable, t.CharClassTable, dfaModel.AcceptTable, len(this.TokenInfos))
}