package regex

type NFAEdge struct {
	Symbol			rune
	TargetState		*NFAState
	IsEmpty			bool
}

func NewNFAEdge(symbol rune, targetState *NFAState) *NFAEdge {
	return &NFAEdge{
		Symbol:      symbol,
		TargetState: targetState,
		IsEmpty:     false,
	}
}

func NewEmptyNFAEdge(targetState *NFAState) *NFAEdge {
	return &NFAEdge{
		Symbol:      0,
		TargetState: targetState,
		IsEmpty:     true,
	}
}
