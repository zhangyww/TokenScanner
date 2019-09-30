package regex

import (
	"io"
	"strings"
	"unicode"
)

var EOF_TOKEN_INDEX int 	= -2

type Scanner struct {
	Engine				*FiniteAutomationEngine
	LexemeValueBuilder	strings.Builder
	Reader				io.RuneReader
	AcceptTable			[]int
	NextRune			rune
}

func NewScanner(scannerInfo *ScannerInfo) *Scanner {
	return &Scanner{
		Engine:             NewFiniteAutomationEngine(scannerInfo.TransitionTable, scannerInfo.CharClassTable),
		LexemeValueBuilder: strings.Builder{},
		Reader:             nil,
		AcceptTable:        scannerInfo.AcceptTable,
		NextRune:           -1,
	}
}

func (this *Scanner) SetReader(reader io.RuneReader) {
	this.Reader = reader
}

func (this *Scanner) Read() *Lexeme {
	if this.Reader == nil {
		return NewLexeme(-1, "")
	}

	this.Engine.Reset()
	this.LexemeValueBuilder.Reset()
	var c rune
	var err error

	if this.NextRune == -1 {
		c, _, err = this.Reader.ReadRune()
		if err != nil {
			return NewLexeme(EOF_TOKEN_INDEX, "")
		}
	} else {
		c = this.NextRune
	}

	//直到第一个非空白符
	for {
		if !unicode.IsSpace(c) {
			break
		}

		c, _, err = this.Reader.ReadRune()
		if err != nil {
			this.NextRune = -1
			return NewLexeme(EOF_TOKEN_INDEX, "")
		}
	}

	this.Engine.Input(c)
	this.LexemeValueBuilder.WriteRune(c)

	if this.Engine.IsAtStoppedState() {
		this.NextRune = -1
		return NewLexeme(-1, this.LexemeValueBuilder.String())
	}

	for {
		c, _, err = this.Reader.ReadRune()
		if err != nil {
			this.NextRune = -1
			return NewLexeme(this.AcceptTable[this.Engine.CurrentState], this.LexemeValueBuilder.String())
		}

		this.Engine.Input(c)

		if this.Engine.IsAtStoppedState() {
			this.NextRune = c
			return NewLexeme(this.AcceptTable[this.Engine.PrevState], this.LexemeValueBuilder.String())
		}
		this.LexemeValueBuilder.WriteRune(c)

	}
}