package regex

type AlternationRegex struct {
	Exp1 IRegex
	Exp2 IRegex
}

func NewAlternationRegex(exp1 IRegex, exp2 IRegex) *AlternationRegex {
	return &AlternationRegex{
		Exp1: exp1,
	 	Exp2: exp2,
	}
}

func (this *AlternationRegex) GetType() RegexType {
	return REGEX_ALTERNATION
}


func (this *AlternationRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertAlternation(this)
}

func (this *AlternationRegex) GetUncompactableCharset() *RuneSet {
	runeSet := this.Exp1.GetUncompactableCharset()
	runeSet.Union(this.Exp2.GetUncompactableCharset())

	return runeSet
}

func (this *AlternationRegex)Many() IRegex {
	return NewKleeneStarRegex(this)
}

func (this *AlternationRegex)Concat(exp IRegex) IRegex {
	return NewConcatenationRegex(this, exp)
}

func (this *AlternationRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *AlternationRegex)Many1() IRegex {
	return this.Concat(this.Many())
}

func (this *AlternationRegex)Optional() IRegex {
	return this.Union(Empty())
}

func (this *AlternationRegex)Repeat(n int) IRegex {
	if n <= 0 {
		return Empty()
	}
	retExp := IRegex(this)
	for idx := 1; idx < n; idx++ {
		retExp = retExp.Concat(this)
	}
	return retExp
}