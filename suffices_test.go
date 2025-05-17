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
	var sis SISuffix
	var d uint8
	v :="100G"
	fmt.Sscanf(v,"%d%v",&d,&sis) // scans d as uint8 so no dp and only up to 255 
	fmt.Printf("%d%s == %+k\n",d,sis,NewSI(float64(d)/sis.Suffix.Scale(sSI)))
	// Output:
	// 100G == 100 000 000k
}



