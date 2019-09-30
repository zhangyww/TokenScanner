package regex

type KleeneStarRegex struct {
	InnerExp IRegex
}

func NewKleeneStarRegex(exp IRegex) *KleeneStarRegex {
	return &KleeneStarRegex{
		InnerExp: exp,
	}
}

func (this *KleeneStarRegex) GetType() RegexType {
	return REGEX_KLEENESTART
}

func (this *KleeneStarRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertKleeneStar(this)
}

func (this *KleeneStarRegex) GetUncompactableCharset() *RuneSet {
	return this.InnerExp.GetUncompactableCharset()
}

func (this *KleeneStarRegex)Many() IRegex {
	if this.GetType() == REGEX_KLEENESTART {
		return this
	}
	return NewKleeneStarRegex(this)
}

func (this *KleeneStarRegex)Concat(exp IRegex) IRegex {
	return NewConcatenationRegex(this, exp)
}

func (this *KleeneStarRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *KleeneStarRegex)Many1() IRegex {
	return this.Concat(this.Many())
}

func (this *KleeneStarRegex)Optional() IRegex {
	return this.Union(Empty())
}

func (this *KleeneStarRegex)Repeat(n int) IRegex {
	if n <= 0 {
		return Empty()
	}
	retExp := IRegex(this)
	for idx := 1; idx < n; idx ++ {
		retExp = retExp.Concat(this)
	}
	return retExp
}