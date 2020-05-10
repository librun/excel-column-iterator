package iterator

import "strconv"

// Alphabet struct iterator
type Alphabet struct {
	startColumn  int
	currentIndex int
	countColumn  int
	row          int
}

// NewIterator new iterator
func NewIterator(from string, len, row int) *Alphabet {
	return &Alphabet{
		startColumn:  ConvertStringToInt(from),
		currentIndex: -1,
		countColumn:  len,
		row:          row,
	}
}

// Next func for block while
func (a *Alphabet) Next() bool {
	a.currentIndex++

	return a.currentIndex < a.countColumn
}

func (a *Alphabet) GetAddress() (int, string) {
	return a.currentIndex, ConvertIntToString(a.startColumn+a.currentIndex) + strconv.Itoa(a.row)
}
