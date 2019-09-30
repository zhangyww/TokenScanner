package regex

type 	RegexType	int

const (
	REGEX_EMPTY 		RegexType = iota
	REGEX_SYMBOL		RegexType = iota
	REGEX_ALTERNATION	RegexType = iota
	REGEX_CONCATENATION	RegexType = iota
	REGEX_KLEENESTART	RegexType = iota

	REGEX_STRINGLITERAL	RegexType = iota
	REGEX_ALTERNATION_CHARSET	RegexType = iota
)

type IRegex interface {
	GetType() RegexType
	Accept(converter *NFAConverter) *NFAModel
	GetUncompactableCharset() *RuneSet

	Many() IRegex
	Concat(exp IRegex) IRegex
	Union(exp IRegex) IRegex
	Many1() IRegex
	Optional() IRegex
	Repeat(n int) IRegex
}

