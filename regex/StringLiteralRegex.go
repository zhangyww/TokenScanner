package regex

type StringLiteralRegex struct {
	Literal		string
}

func NewStringLiteralRegex(str string) *StringLiteralRegex {
	return &StringLiteralRegex{
		Literal: str,
	}
}

func (this *StringLiteralRegex) GetType() RegexType {
	return REGEX_STRINGLITERAL
}

func (this *StringLiteralRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertStringLiteral(this)
}

func (this *StringLiteralRegex) GetUncompactableCharset() *RuneSet {
	runeSet := NewRuneSet()

	for _, c := range []rune(this.Literal) {
		runeSet.Add(c)
	}

	return runeSet
}

func (this *StringLiteralRegex)Many() IRegex {
	return NewKleeneStarRegex(this)
}

func (this *StringLiteralRegex)Concat(exp IRegex) IRegex {
	return NewConcatenationRegex(this, exp)
}

func (this *StringLiteralRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *StringLiteralRegex)Many1() IRegex {
	return this.Concat(this.Many())
}

func (this *StringLiteralRegex)Optional() IRegex {
	return this.Union(Empty())
}

func (this *StringLiteralRegex)Repeat(n int) IRegex {
	if n <= 0 {
		return Empty()
	}
	retExp := IRegex(this)
	for idx := 1; idx < n; idx ++ {
		retExp = retExp.Concat(this)
	}
	return retExp
}