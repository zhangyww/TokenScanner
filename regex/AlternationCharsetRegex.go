package regex

type AlternationCharsetRegex struct {
	Charset		[]rune
}

func NewAlternationCharsetRegex(charset []rune) *AlternationCharsetRegex {
	return &AlternationCharsetRegex{
		Charset: charset,
	}
}

func (this *AlternationCharsetRegex) GetType() RegexType {
	return REGEX_ALTERNATION_CHARSET
}

func (this *AlternationCharsetRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertAlternationCharset(this)
}

func (this *AlternationCharsetRegex) GetUncompactableCharset() *RuneSet {
	return NewRuneSet()
}

func (this *AlternationCharsetRegex)Many() IRegex {
	return NewKleeneStarRegex(this)
}

func (this *AlternationCharsetRegex)Concat(exp IRegex) IRegex {
	return NewConcatenationRegex(this, exp)
}

func (this *AlternationCharsetRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *AlternationCharsetRegex)Many1() IRegex {
	return this.Concat(this.Many())
}

func (this *AlternationCharsetRegex)Optional() IRegex {
	return this.Union(Empty())
}

func (this *AlternationCharsetRegex)Repeat(n int) IRegex {
	if n <= 0 {
		return Empty()
	}
	retExp := IRegex(this)
	for idx := 1; idx < n; idx ++ {
		retExp = retExp.Concat(this)
	}
	return retExp
}