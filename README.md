# TokenScanner

(装配脑袋)[https://www.cnblogs.com/ninputer/] 的VBF库中Scanner的go实现。

# Example

```go
lexicon := regex.NewLexicon()
lexer := lexicon.DefaultLexer

IF := lexer.DefineToken(regex.Literal("if"))
ELSE := lexer.DefineToken(regex.Literal("else"))
ID := lexer.DefineToken(regex.Range('a','z').Concat(
	(regex.Range('a','z').Union(regex.Range('0','9')).Many())))
NUM := lexer.DefineToken(regex.Range('0','9').Many1())

scannerInfo := lexicon.CreateScannerInfo()
scanner := regex.NewScanner(scannerInfo)

source := "asdf04a 1107 else if"
scanner.SetReader(bytes.NewReader([]byte(source)))

l1 := scanner.Read()
fmt.Println("l1.Token == ", l1.TokenIndex)
fmt.Println("l1.Value == ", l1.Value)
fmt.Println("ID.INdex == ", ID.Index)
fmt.Println("----------------------------")

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

l4 := scanner.Read()
fmt.Println("l4.Token == ", l4.TokenIndex)
fmt.Println("l4.Value == ", l4.Value)
fmt.Println("EOFTOKEN.INdex == ", regex.EOF_TOKEN_INDEX)
fmt.Println("----------------------------")

```

输出结果为

```
l1.Token ==  2
l1.Value ==  asdf04a
ID.INdex ==  2
----------------------------
l2.Token ==  3
l2.Value ==  1107
NUM.INdex ==  3
----------------------------
l3.Token ==  1
l3.Value ==  else
ELSE.INdex ==  1
----------------------------
l3.Token ==  0
l3.Value ==  if
IF.INdex ==  0
----------------------------
l4.Token ==  -2
l4.Value ==  
EOFTOKEN.INdex ==  -2
----------------------------
```