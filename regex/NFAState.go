package regex

type NFAState struct {
	OutEdges	[]*NFAEdge
	Index		int
	TokenIndex	int
}

func NewNFAState() *NFAState {
	return &NFAState{
		OutEdges:   []*NFAEdge{},
		Index:      0,
		TokenIndex: -1,
	}
}

func (this *NFAState) AddEdge(edge *NFAEdge) {
	this.OutEdges = append(this.OutEdges, edge)
}

func (this *NFAState) AddEmptyEdgeTo(targetState *NFAState) {
	this.OutEdges = append(this.OutEdges, NewEmptyNFAEdge(targetState))
}
