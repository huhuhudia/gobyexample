package main

import "fmt"
import "os"

type Point struct{
	X, Y int
}

func main(){
	p := Point{1,2}
	// print the value of p 
	fmt.Printf("%v\n",p)
	// print the value and its field name
	fmt.Printf("%+v\n",p)

	fmt.Printf("%#v\n",p)

	// type
	fmt.Printf("%T\n",p)
	// format boolean
	fmt.Printf("%t\n",true)

	// binary represent
	fmt.Printf("%b\n",p)

	// print the character corresponding to the given integer
	fmt.Printf("%c\n",14)

	// print the hex encoding 
	fmt.Printf("%x\n",16)

	// print float
	fmt.Printf("%x\n",16)
	//string formating
	fmt.Printf("%s\n","123")
	// double qoute string 
	fmt.Printf("%q\n","\"double quote \"")

	// hex this 
	fmt.Printf("%x\n","\"double quote \"")

	// pointer representation
	fmt.Printf("%p\n", &p)

	// specify the width 
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// - to the left
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 12.1, 345.2)

	// format and print to io.writer
	fmt.Fprintf(os.Stdout, "an %s\n", "error")
}

