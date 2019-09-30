package regex

type ScannerInfo struct {
	EOFTokenIndex		int
	AcceptTable			[]int
	CharClassTable		[]rune
	TransitionTable		[][]rune
	TokenCount			int
}

func NewSannerInfo(transitionTable [][]rune, charCLassTable []rune, acceptTable []int, tokenCount int) *ScannerInfo{
	return &ScannerInfo{
		EOFTokenIndex:   tokenCount,
		AcceptTable:     acceptTable,
		CharClassTable:  charCLassTable,
		TransitionTable: transitionTable,
		TokenCount:      tokenCount,
	}
}
