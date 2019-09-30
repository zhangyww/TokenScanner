package regex

type NFAModel struct {
	States			[]*NFAState
	TailState		*NFAState
	EntryEdge		*NFAEdge
}

func NewNFAModel() *NFAModel{
	return &NFAModel{
		States:      []*NFAState{},
		TailState: nil,
		EntryEdge:   nil,
	}
}

func (this *NFAModel) AddState(state *NFAState) {
	this.States = append(this.States, state)
	state.Index = len(this.States) - 1
}

func (this *NFAModel) AddStates(states []*NFAState) {

	for idx, _ := range states {
		this.AddState(states[idx])
	}
	//for e:=states.Front(); e!=nil; e = e.Next() {
	//	this.AddState(e.Value.(*NFAState))
	//}
}