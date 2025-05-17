package scalingSuffix

import "fmt"

func ExampleNewSISuffix(){
	fmt.Printf("%q\n",NewSISuffix('1'))
	fmt.Printf("%q\n",NewSISuffix(' '))
	fmt.Printf("%q\n",NewSISuffix('k'))
	fmt.Printf("%q\n",NewSISuffix('P'))
	// Output:
	// ""
	// "" 
	// "k"
	// "P"
}

func ExampleSISuffix_Scan(){
	var s SISuffix
	var d uint8
	fmt.Sscanf("100G","%d%v",&d,&s)
	fmt.Printf("%d%s\n",d,s)
	// Output:
	// 100G

}



