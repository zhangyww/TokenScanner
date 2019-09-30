package regex

import (
	"strings"
)

type CompressedTransitionTable struct {
	CharClassTable			[]rune
	StateSetDict			map[string]rune
	DfaStates				[]*DFAState
	CompactCharsetManager	*CompactCharsetManager
	TransitionTable			[][]rune
}

func NewCompressedTransitionTable(dfaModel *DFAModel) *CompressedTransitionTable {
	return &CompressedTransitionTable{
		CharClassTable:        make([]rune, 65536),
		StateSetDict:          make(map[string]rune),
		DfaStates:             dfaModel.DfaStates,
		CompactCharsetManager: dfaModel.CompactCharsetManager,
		TransitionTable:       make([][]rune, len(dfaModel.DfaStates)),
	}
}

func Compress(dfaModel *DFAModel) *CompressedTransitionTable{
	compressedTransitionTable := NewCompressedTransitionTable(dfaModel)
	compressedTransitionTable.Compress()

	return compressedTransitionTable
}

func (this *CompressedTransitionTable) Compress() {
	transitionTable	:= make([]map[rune]rune, len(this.DfaStates))
	dfaStateCount := len(this.DfaStates)

	for idx := 0; idx < dfaStateCount; idx ++ {
		transitionTable[idx] = make(map[rune]rune)

		for _, edge := range this.DfaStates[idx].OutEdges {
			transitionTable[idx][edge.CharsetClass] = rune(edge.TargetState.Index)
		}
	}

	transitionColumnTable := [][]rune{}
	charClassToCharMapTable := this.CompactCharsetManager.CreateCharClassToCharMapTable()

	for charClass := rune(0); charClass <= this.CompactCharsetManager.MaxIndex; charClass ++ {
		columnSequence := []rune{}
		for _, rowMap := range transitionTable {
			columnSequence = append(columnSequence, rowMap[charClass])
		}

		builder := strings.Builder{}
		for _, c := range columnSequence {
			builder.WriteRune(c)
			builder.WriteRune(',')
		}
		columnSignature := builder.String()

		if _, isExist := this.StateSetDict[columnSignature]; isExist {
			newCharClass := this.StateSetDict[columnSignature]
			for charSymbol, _ := range charClassToCharMapTable[charClass].Data {
				this.CharClassTable[charSymbol] = newCharClass
			}
		} else {
			nextIndex := len(transitionColumnTable)

			transitionColumnTable = append(transitionColumnTable, columnSequence)

			this.StateSetDict[columnSignature] = rune(nextIndex)

			for charSymbol, _ := range charClassToCharMapTable[charClass].Data {
					this.CharClassTable[charSymbol] = rune(nextIndex)
			}
		}
	}

	invalidColumn := make([]rune, len(this.DfaStates))
	invalidIndex := len(transitionColumnTable)
	transitionColumnTable = append(transitionColumnTable, invalidColumn)

	for charSymbol, _ := range charClassToCharMapTable[0].Data {
		this.CharClassTable[charSymbol] = rune(invalidIndex)
	}

	//生成最终的压缩转换表
	for rowIndex := 0; rowIndex < dfaStateCount; rowIndex ++ {
		this.TransitionTable[rowIndex] = make([]rune, len(transitionColumnTable))
		for columnIndex, column := range transitionColumnTable {
			this.TransitionTable[rowIndex][columnIndex] = column[rowIndex]
		}
	}

}