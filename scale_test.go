package scalingSuffix

import "fmt"
import "strings"
import "strconv"



func ExamplePrint(){
	f,e,_:=strings.Cut(fmt.Sprintf("%e",123_456_789.0),"e")
	ev,_:=strconv.Atoi(e)
	ei,s:=SmallerSuffix(ev)
	fmt.Printf("%v%c\n",MovePointRight(f,ev-ei),s)
	f,e,_=strings.Cut(fmt.Sprintf("%e",0.02),"e")
	ev,_=strconv.Atoi(e)
	ei,s=SmallerSuffix(ev)
	fmt.Printf("%v%c\n",TrimNumber(MovePointRight(f,ev-ei)),s)
	// Output:
	// 123.4568M
	// 2c
}

