package regex

type RuneSet struct {
	Data map[rune]struct{}
}

func NewRuneSet() *RuneSet {
	return &RuneSet{
		Data: make(map[rune]struct{}),
	}
}

func (this *RuneSet) Count() int {
	return len(this.Data)
}

func (this *RuneSet) Add(c rune){
	this.Data[c] = struct{}{}
}

func (this *RuneSet) Union(other *RuneSet) {
	for c, _ := range other.Data {
		this.Data[c] = struct{}{}
	}
}

func (this *RuneSet) Intersect(other *RuneSet) {
	retData := make(map[rune]struct{})
	for c, _ := range this.Data {
		if _, has := other.Data[c]; has {
			retData[c] = struct{}{}
		}
	}
	this.Data = retData
}

func (this *RuneSet) Equals(other *RuneSet) bool {
	equalCount := 0
	for c, _ := range this.Data {
		_, isExist := other.Data[c]
		if isExist {
			equalCount++
		} else {
			return false
		}
	}
	return equalCount == other.Count()
}