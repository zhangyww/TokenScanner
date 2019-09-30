package regex

import (
	"sort"
)

type DFAModel struct {
	AcceptTable		[]int
	DfaStates		[]*DFAState
	Lexicon			*Lexicon
	NfaModel		*NFAModel
	CompactCharsetManager 	*CompactCharsetManager
}

func NewDFAModel(lexicon *Lexicon) *DFAModel {
	return &DFAModel{
		AcceptTable: []int{},
		DfaStates:   []*DFAState{},
		Lexicon:     lexicon,
		NfaModel:    nil,
	}
}

func CreateDFAModel(lexicon *Lexicon) *DFAModel {
	dfaModel := NewDFAModel(lexicon)
	dfaModel.ConvertLexcionToNFA()
	dfaModel.ConvertNFAToDFA()

	return dfaModel
}

func (this *DFAModel)ConvertLexcionToNFA() {
	this.CompactCharsetManager = this.Lexicon.CreateCompactCharsetManager()
	nfaConverter := NewNFAConverter(this.CompactCharsetManager)

	entryState := NewNFAState()
	nfaModel := NewNFAModel()

	nfaModel.AddState(entryState)
	tokenInfos := this.Lexicon.TokenInfos

	for idx, _ := range tokenInfos {
		tokenNFAModel := tokenInfos[idx].CreateFiniteAutomatonModel(nfaConverter)

		entryState.AddEdge(tokenNFAModel.EntryEdge)
		nfaModel.AddStates(tokenNFAModel.States)
	}

	nfaModel.EntryEdge = NewEmptyNFAEdge(entryState)

	this.NfaModel = nfaModel
}

func (this *DFAModel) ConvertNFAToDFA() {
	//nfaStates := this.NfaModel.States

	state0 := NewDFAState()
	this.AddState(state0)

	preState1 := NewDFAState()
	nfaStartIndex := this.NfaModel.EntryEdge.TargetState.Index
	preState1.NFAStatesIndexes.Add(rune(nfaStartIndex))

	state1 := this.GetClosure(preState1)
	this.AddState(state1)

	dfaStateCountIndex := 1
	curDfaStateIdx := 0
	newStates := make([]*DFAState, this.CompactCharsetManager.MaxIndex + 1)

	for {
		if curDfaStateIdx > dfaStateCountIndex { break }

		sourceState := this.DfaStates[curDfaStateIdx]

		for charsetClass := this.CompactCharsetManager.MinIndex; charsetClass <= this.CompactCharsetManager.MaxIndex; charsetClass++ {
			e := this.GetDFAState(sourceState, charsetClass)
			newStates[charsetClass] = e
		}

		for charsetClass := this.CompactCharsetManager.MinIndex; charsetClass <= this.CompactCharsetManager.MaxIndex; charsetClass++ {
			outClassClosure := newStates[charsetClass]
			isSetExist := false

			for idx := 0; idx <= dfaStateCountIndex; idx++ {
				if outClassClosure.NFAStatesIndexes.Equals(this.DfaStates[idx].NFAStatesIndexes) {
					newEdge := NewDFAEdge(charsetClass, this.DfaStates[idx])
					sourceState.AddEdge(newEdge)

					isSetExist = true
				}
			}

			if !isSetExist {
				dfaStateCountIndex ++
				this.AddState(outClassClosure)

				newEdge := NewDFAEdge(charsetClass, outClassClosure)
				sourceState.AddEdge(newEdge)
			}
		}

		curDfaStateIdx ++
	}

}

func (this *DFAModel) AddState(state *DFAState) {
	this.DfaStates = append(this.DfaStates, state)
	state.Index = len(this.DfaStates) - 1

	this.AcceptTable = append(this.AcceptTable, -1)

	tokenInfos := this.Lexicon.TokenInfos
	nfaStates := this.NfaModel.States

	acceptStates := []*TokenInfo{}

	for nfaIndex, _ := range state.NFAStatesIndexes.Data {
		tokenIndex := nfaStates[nfaIndex].TokenIndex
		if tokenIndex >= 0 {
			acceptStates = append(acceptStates, tokenInfos[tokenIndex])
		}
	}
	sort.Sort(ByTokenIndex(acceptStates))

	if len(acceptStates) > 0 {
		this.AcceptTable[state.Index] = acceptStates[0].Token.Index
	}
}

func (this *DFAModel) GetClosure(state *DFAState) *DFAState {
	closure := NewDFAState()
	closure.NFAStatesIndexes.Union(state.NFAStatesIndexes)

	nfaStates := this.NfaModel.States
	changed := true

	for {
		if !changed { break}

		changed = false
		lastStateSet := make([]int, closure.NFAStatesIndexes.Count())
		idx := 0
		for i, _ := range closure.NFAStatesIndexes.Data {
			lastStateSet[idx] = int(i)
			idx++
		}

		for _, stateIndex := range lastStateSet {
			nfaState := nfaStates[stateIndex]
			outEdges := nfaState.OutEdges
			edgeCount := len(outEdges)
			for idx = 0; idx < edgeCount; idx++ {
				edge := outEdges[idx]

				if edge.IsEmpty {
					target := edge.TargetState

					_, alreadyExist := closure.NFAStatesIndexes.Data[rune(target.Index)]

					closure.NFAStatesIndexes.Add(rune(target.Index))
					changed = alreadyExist || changed
				}
			}
		}
	}

	return closure
}

func (this *DFAModel) GetDFAState(start *DFAState, charsetClass rune) *DFAState {
	target := NewDFAState()
	nfaStates := this.NfaModel.States

	for nfaIndex, _ := range start.NFAStatesIndexes.Data {
		outEdges := nfaStates[nfaIndex].OutEdges
		edgeCount := len(outEdges)

		for idx := 0; idx < edgeCount; idx++ {
			edge := outEdges[idx]
			if !edge.IsEmpty && charsetClass == edge.Symbol {
				target.NFAStatesIndexes.Add(rune(edge.TargetState.Index))
			}
		}
	}

	return this.GetClosure(target)
}