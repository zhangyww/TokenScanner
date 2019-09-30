package regex



type EmptyRegex struct {

}

func (this *EmptyRegex) GetType() RegexType {
	return REGEX_EMPTY
}

func (this *EmptyRegex) Accept(converter *NFAConverter) *NFAModel {
	return converter.ConvertEmpty(this)
}

func (this *EmptyRegex) GetUncompactableCharset() *RuneSet {
	return NewRuneSet()
}

func (this *EmptyRegex)Many() IRegex {
	return this
}

func (this *EmptyRegex)Concat(exp IRegex) IRegex {
	return exp
}

func (this *EmptyRegex)Union(exp IRegex) IRegex {
	if this == exp {
		return this
	}
	return NewAlternationRegex(this, exp)
}

func (this *EmptyRegex)Many1() IRegex {
	return this
}

func (this *EmptyRegex)Optional() IRegex {
	return this
}

func (this *EmptyRegex)Repeat(n int) IRegex {
	return this
}