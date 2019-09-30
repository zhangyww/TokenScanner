package regex

type DFAEdge struct {
	CharsetClass		rune
	TargetState			*DFAState
}

func NewDFAEdge(charsetClass rune, targetState *DFAState) *DFAEdge {
	return &DFAEdge{
		CharsetClass: charsetClass,
		TargetState:  targetState,
	}
}
