package main

import(
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos/1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
	NextYear := time.Date(2020,1,1,0,0,0,0,time.UTC)
	LastYear := time.Date(2019,1,1,0,0,0,0,time.UTC)
	allYear := NextYear.Sub(LastYear)
	diff := NextYear.Sub(now)
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Seconds()/allYear.Seconds())
}