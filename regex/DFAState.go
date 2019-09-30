package regex

type DFAState struct {
	NFAStatesIndexes	*RuneSet
	OutEdges			[]*DFAEdge
	Index				int
}

func NewDFAState() *DFAState {
	return &DFAState{
		NFAStatesIndexes: NewRuneSet(),
		OutEdges:         []*DFAEdge{},
		Index:            0,
	}
}

func (this *DFAState) AddEdge (edge *DFAEdge) {
	this.OutEdges = append(this.OutEdges, edge)
}