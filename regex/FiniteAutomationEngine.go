package regex

type FiniteAutomationEngine struct {
	TransitionTable		[][]rune
	CharClassTable		[]rune
	CurrentState		rune
	PrevState			rune
}

func NewFiniteAutomationEngine(transitionTable [][]rune, charClassTable []rune) *FiniteAutomationEngine {
	return &FiniteAutomationEngine{
		TransitionTable: transitionTable,
		CharClassTable:  charClassTable,
		CurrentState:    1,
		PrevState:       1,
	}
}

func (this *FiniteAutomationEngine) IsAtStoppedState() bool {
	return this.CurrentState == 0
}

func (this *FiniteAutomationEngine) Reset() {
	this.CurrentState = 1
	this.PrevState = 1
}

func (this *FiniteAutomationEngine) Input(c rune) {
	charClass := this.CharClassTable[c]
	nextState := this.TransitionTable[this.CurrentState][charClass]

	this.PrevState = this.CurrentState
	this.CurrentState = nextState
}

func (this *FiniteAutomationEngine) InputString(str string) {
	for _, c := range []rune(str) {
		this.Input(c)
	}
}
