package main

import (
	"time"
	"fmt"
)

func main(){
	print(time.Now().String())
	timer1 := time.NewTimer(2*time.Second)
	<-timer1.C
	fmt.Println("Timer1 expired")

	time2 := time.NewTimer(time.Second)

	go func(){
		<- time2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := time2.Stop()  // cancel the timer2 sleep
	if stop2{
		print(time.Now().String())
		fmt.Println("timer 2 stopped ")
	}
}