package regex

type CompactCharsetManager struct {
	MinIndex 		rune
	MaxIndex		rune
	CharClassTable	[]rune
}

func NewCompatCharsetManager(charClassTable []rune, maxIndex rune) *CompactCharsetManager {
	this := &CompactCharsetManager{
		MinIndex:       1,
		MaxIndex:       maxIndex,
		CharClassTable: charClassTable,
	}

	return this
}

func (this *CompactCharsetManager) GetCompactClass(c rune) rune {
	cls := this.CharClassTable[c]

	return cls
}


func (this *CompactCharsetManager) CreateCharClassToCharMapTable() []*RuneSet {
	result := make([]*RuneSet, this.MaxIndex + 1)
	for idx := rune(0); idx <= this.MaxIndex; idx ++ {
		result[idx] = NewRuneSet()
	}

	for idx := rune(0); idx <= rune(65535); idx ++ {
		charClass := this.CharClassTable[idx]
		result[charClass].Add(idx)
	}

	return result
}
