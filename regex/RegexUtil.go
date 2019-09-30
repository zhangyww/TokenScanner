package regex

var emptyInstance *EmptyRegex = &EmptyRegex{}

func Empty() IRegex {
	return emptyInstance
}

func Symbol(c rune) IRegex {
	return NewSymbolRegex(c)
}

func Literal(str string) IRegex {
	return NewStringLiteralRegex(str)
}

func Range(min rune, max rune) IRegex {
	count := (max - min + 1)
	if count <= 0 {
		return Empty()
	} else if count == 1 {
		return Symbol(min)
	}

	var exp IRegex = Symbol(min)

	for c := min + 1; c <= max; c ++ {
		exp = exp.Union(Symbol(c))
	}

	return exp
}