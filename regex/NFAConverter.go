package regex

type NFAConverter struct {
	CompactCharsetManager 	*CompactCharsetManager
}

func NewNFAConverter(compactCharsetManager *CompactCharsetManager) *NFAConverter{
	return &NFAConverter{
		CompactCharsetManager: compactCharsetManager,
	}
}

func (this *NFAConverter) Convert(exp IRegex) *NFAModel {
	return exp.Accept(this)
}

func (this *NFAConverter) ConvertEmpty(exp *EmptyRegex) *NFAModel {
	tailState := NewNFAState()
	entryEdge := NewEmptyNFAEdge(tailState)

	emptyNFA := NewNFAModel()

	emptyNFA.AddState(tailState)
	emptyNFA.TailState = tailState
	emptyNFA.EntryEdge = entryEdge

	return emptyNFA
}

func (this *NFAConverter) ConvertSymbol(exp *SymbolRegex) *NFAModel {

	tailState := NewNFAState()
	//这里可能需要一个转换，将exp.Symbol转成一个数值
	charsetClass := this.CompactCharsetManager.GetCompactClass(exp.Symbol)

	entryEdge := NewNFAEdge(charsetClass, tailState)

	symbolNFA := NewNFAModel()

	symbolNFA.AddState(tailState)
	symbolNFA.TailState = tailState
	symbolNFA.EntryEdge = entryEdge

	return symbolNFA
}

func (this *NFAConverter) ConvertAlternation(exp *AlternationRegex) *NFAModel {
	nfa1 := this.Convert(exp.Exp1)
	nfa2 := this.Convert(exp.Exp2)

	headState := NewNFAState()
	tailState := NewNFAState()

	headState.AddEdge(nfa1.EntryEdge)
	headState.AddEdge(nfa2.EntryEdge)

	nfa1.TailState.AddEmptyEdgeTo(tailState)
	nfa2.TailState.AddEmptyEdgeTo(tailState)

	alternationNFA := NewNFAModel()

	alternationNFA.AddState(headState)
	alternationNFA.AddStates(nfa1.States)
	alternationNFA.AddStates(nfa2.States)
	alternationNFA.AddState(tailState)

	alternationNFA.TailState = tailState
	alternationNFA.EntryEdge = NewEmptyNFAEdge(headState)

	return alternationNFA
}

func (this *NFAConverter) ConvertConcatenation(exp *ConcatenationRegex) *NFAModel {
	leftNFA := this.Convert(exp.LeftExp)
	rightNFA := this.Convert(exp.RightExp)

	leftNFA.TailState.AddEdge(rightNFA.EntryEdge)

	concatenationNFA := NewNFAModel()
	concatenationNFA.AddStates(leftNFA.States)
	concatenationNFA.AddStates(rightNFA.States)

	concatenationNFA.TailState = rightNFA.TailState
	concatenationNFA.EntryEdge = leftNFA.EntryEdge

	return concatenationNFA
}

func (this *NFAConverter) ConvertKleeneStar(exp *KleeneStarRegex) *NFAModel {
	innerNFA := this.Convert(exp.InnerExp)

	tailState := NewNFAState()
	entryEdge := NewEmptyNFAEdge(tailState)

	innerNFA.TailState.AddEmptyEdgeTo(tailState)
	tailState.AddEdge(innerNFA.EntryEdge)

	kleeneStarNFA := NewNFAModel()
	kleeneStarNFA.AddStates(innerNFA.States)
	kleeneStarNFA.AddState(tailState)

	kleeneStarNFA.TailState = tailState
	kleeneStarNFA.EntryEdge = entryEdge

	return kleeneStarNFA
}

func (this *NFAConverter) ConvertStringLiteral(exp *StringLiteralRegex) *NFAModel {
	literalNFA := NewNFAModel()

	var lastState *NFAState = nil

	for _, c := range []rune(exp.Literal) {
		symbolState := NewNFAState()
		charsetClass := this.CompactCharsetManager.GetCompactClass(c)
		symbolEdge := NewNFAEdge(charsetClass, symbolState)

		if lastState != nil {
			lastState.AddEdge(symbolEdge)
		} else {
			literalNFA.EntryEdge = symbolEdge
		}
		lastState = symbolState
		literalNFA.AddState(symbolState)
	}
	literalNFA.TailState = lastState

	return literalNFA
}

func (this *NFAConverter) ConvertAlternationCharset(exp *AlternationCharsetRegex) *NFAModel {
	headState := NewNFAState()
	tailState := NewNFAState()

	charsetNFA := NewNFAModel()
	charsetNFA.AddState(headState)

	for _, c := range exp.Charset {
		charsetClass := this.CompactCharsetManager.GetCompactClass(c)
		headState.AddEdge(NewNFAEdge(charsetClass, tailState))
	}
	charsetNFA.AddState(tailState)

	charsetNFA.EntryEdge = NewEmptyNFAEdge(headState)
	charsetNFA.TailState = tailState

	return charsetNFA
}