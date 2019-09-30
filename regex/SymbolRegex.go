package regex

type SymbolRegex struct {
	Symbol	rune
}

func NewSymbolRegex(symbol rune) *SymbolRegex {
	return &SymbolRegex{
		Symbol: symbol,
	}
}

func (this *SymbolRegex) GetType() RegexType {
	return REGEX_SYMBOL
}

func (this *SymbolRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertSymbol(this)
}

func (this *SymbolRegex) GetUncompactableCharset() *RuneSet {
	runeSet := NewRuneSet()
	runeSet.Add(this.Symbol)

	return runeSet
}

func (this *SymbolRegex)Many() IRegex {
	return NewKleeneStarRegex(this)
}

func (this *SymbolRegex)Concat(exp IRegex) IRegex {
	return NewConcatenationRegex(this, exp)
}

func (this *SymbolRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *SymbolRegex)Many1() IRegex {
	return this.Concat(this.Many())
}

func (this *SymbolRegex)Optional() IRegex {
	return this.Union(Empty())
}

func (this *SymbolRegex)Repeat(n int) IRegex {
	if n <= 0 {
		return Empty()
	}
	retExp := IRegex(this)
	for idx := 1; idx < n; idx ++ {
		retExp = retExp.Concat(this)
	}
	return retExp
}