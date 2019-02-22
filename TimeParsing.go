package main

import (
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	if e != nil{
		fmt.Println(e)
	}
	p(t1)

	p(t.Format("3:04 PM"))

}