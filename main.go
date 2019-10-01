package main

import (
	"bytes"
	"fmt"
	"github.com/zhangyww/TokenScanner/regex"
)

func main() {
	lexicon := regex.NewLexicon()
	lexer := lexicon.DefaultLexer

	IF := lexer.DefineToken(regex.Literal("if"))
	ELSE := lexer.DefineToken(regex.Literal("else"))
	ID := lexer.DefineToken(regex.Range('a','z').Concat(
		(regex.Range('a','z').Union(regex.Range('0','9')).Many())))
	NUM := lexer.DefineToken(regex.Range('0','9').Many1())
	//AORB := lexer.DefineToken(regex.NewAlternationRegex(regex.NewAlternationRegex(regex.Symbol('a'),regex.Symbol('b')),regex.Symbol('c')))

	scannerInfo := lexicon.CreateScannerInfo()
	scanner := regex.NewScanner(scannerInfo)

	source := "asdf04a 1107 else if"
	scanner.SetReader(bytes.NewReader([]byte(source)))

	l1 := scanner.Read()
	fmt.Println("l1.Token == ", l1.TokenIndex)
	fmt.Println("l1.Value == ", l1.Value)
	fmt.Println("ID.INdex == ", ID.Index)
	fmt.Println("----------------------------")

	//l1 = scanner.Read()
	//fmt.Println("l1.Token == ", l1.TokenIndex)
	//fmt.Println("l1.Value == ", l1.Value)
	//fmt.Println("ID.INdex == ", ID.Index)
	//
	//l1 = scanner.Read()
	//fmt.Println("l1.Token == ", l1.TokenIndex)
	//fmt.Println("l1.Value == ", l1.Value)
	//fmt.Println("ID.INdex == ", ID.Index)
	//
	//l1 = scanner.Read()
	//fmt.Println("l1.Token == ", l1.TokenIndex)
	//fmt.Println("l1.Value == ", l1.Value)
	//fmt.Println("ID.INdex == ", ID.Index)

	l2 := scanner.Read()
	fmt.Println("l2.Token == ", l2.TokenIndex)
	fmt.Println("l2.Value == ", l2.Value)
	fmt.Println("NUM.INdex == ", NUM.Index)
	fmt.Println("----------------------------")

	l3 := scanner.Read()
	fmt.Println("l3.Token == ", l3.TokenIndex)
	fmt.Println("l3.Value == ", l3.Value)
	fmt.Println("ELSE.INdex == ", ELSE.Index)
	fmt.Println("----------------------------")

	l3 = scanner.Read()
	fmt.Println("l3.Token == ", l3.TokenIndex)
	fmt.Println("l3.Value == ", l3.Value)
	fmt.Println("IF.INdex == ", IF.Index)
	fmt.Println("----------------------------")

	//
	//fmt.Println("IF.INdex == ", IF.Index)


	l4 := scanner.Read()
	fmt.Println("l4.Token == ", l4.TokenIndex)
	fmt.Println("l4.Value == ", l4.Value)
	fmt.Println("EOFTOKEN.INdex == ", regex.EOF_TOKEN_INDEX)
	fmt.Println("----------------------------")

	//dfaModel := regex.CreateDFAModel(lexicon)
	//transitionTable := regex.Compress(dfaModel)
	//
	//fmt.Println("ifToken %V", ifToken)
	//fmt.Println("elseToken %V", elseToken)
	//fmt.Println("whileToken %V", whileToken)
	//fmt.Println("forToken %V", forToken)
	//
	//fmt.Println("transitionTable %V", transitionTable)
	//fmt.Println("i == ",  transitionTable.CharClassTable['i'])
	//fmt.Println("f == ",  transitionTable.CharClassTable['f'])
	//fmt.Println("e == ",  transitionTable.CharClassTable['e'])
	//fmt.Println("l == ",  transitionTable.CharClassTable['l'])
	//fmt.Println("s == ",  transitionTable.CharClassTable['s'])
	//fmt.Println("w == ",  transitionTable.CharClassTable['w'])
	//fmt.Println("h == ",  transitionTable.CharClassTable['h'])
	//fmt.Println("f == ",  transitionTable.CharClassTable['f'])
	//fmt.Println("o == ",  transitionTable.CharClassTable['o'])
	//fmt.Println("r == ",  transitionTable.CharClassTable['r'])
	//fmt.Println("******************************")
	//fmt.Println(transitionTable.CharClassTable[:256])
	//fmt.Println("******************************")
	//for idx, t := range transitionTable.TransitionTable {
	//	fmt.Printf("state[%d] >>>> ", idx)
	//	fmt.Println(t)
	//}
	//fmt.Println("******************************")
	//
	//fmt.Println(dfaModel.AcceptTable)

}
