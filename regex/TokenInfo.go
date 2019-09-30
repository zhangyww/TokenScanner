package regex

type ByTokenIndex []*TokenInfo

func (a ByTokenIndex) Len() int { return len(a) }
func (a ByTokenIndex) Swap(i int, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTokenIndex) Less(i int, j int) bool { return a[i].Token.Index < a[j].Token.Index }

type TokenInfo struct {
	Lexicon *Lexicon
	Lexer   *Lexer
	Token   *Token
	Exp     IRegex
}

func NewTokenInfo(exp IRegex, lexicon *Lexicon, lexer *Lexer, token *Token) *TokenInfo {
	return &TokenInfo{
		Lexicon: lexicon,
		Lexer:   lexer,
		Token:   token,
		Exp:     exp,
	}
}

func (this *TokenInfo) CreateFiniteAutomatonModel(converter *NFAConverter) *NFAModel {
	nfaModel := converter.Convert(this.Exp)

	nfaModel.TailState.TokenIndex = this.Token.Index

	return nfaModel
}

