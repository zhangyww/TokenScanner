package regex

type ConcatenationRegex struct {
	LeftExp  IRegex
	RightExp IRegex
}

func NewConcatenationRegex(leftExp IRegex, rightRegex IRegex) *ConcatenationRegex {
	return &ConcatenationRegex{
		LeftExp:  leftExp,
		RightExp: rightRegex,
	}
}

func (this *ConcatenationRegex) GetType() RegexType {
	return REGEX_CONCATENATION
}

func (this *ConcatenationRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertConcatenation(this)
}

func (this *ConcatenationRegex) GetUncompactableCharset() *RuneSet {
	runeSet := this.LeftExp.GetUncompactableCharset()
	runeSet.Union(this.RightExp.GetUncompactableCharset())

	return runeSet
}

func (this *ConcatenationRegex)Many() IRegex {
	return NewKleeneStarRegex(this)
}

func (this *ConcatenationRegex)Concat(exp IRegex) IRegex {
	return NewConcatenationRegex(this, exp)
}

func (this *ConcatenationRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *ConcatenationRegex)Many1() IRegex {
	return this.Concat(this.Many())
}

func (this *ConcatenationRegex)Optional() IRegex {
	return this.Union(Empty())
}

func (this *ConcatenationRegex)Repeat(n int) IRegex {
	if n <= 0 {
		return Empty()
	}
	retExp := IRegex(this)
	for idx := 1; idx < n; idx ++ {
		retExp = retExp.Concat(this)
	}
	return retExp
}